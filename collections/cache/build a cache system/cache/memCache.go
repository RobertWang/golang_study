// cache/memCache.go
package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// 内存缓存结构
type memCache struct {
	// 最大内存大小
	maxMemorySize int64
	// 最大内存字符串表示
	maxMemorySizeStr string
	// 当前内存大小
	currMemorySize int64

	// 缓坡键值对
	values map[string]*memCacheValue

	// 读写锁
	locker sync.RWMutex

	// 清除过期缓存的时间间隔
	cleanExpiredItemTimeInterval time.Duration
}

// 内存缓存值结构
type memCacheValue struct {
	// value 值
	val interface{}
	// 过期时间
	expireTime time.Time
	// 有效时长
	expire time.Duration
	// value 大小，用于计算当前内存大小
	size int64
}

// 轮询清理过期 key
func (mc *memCache) cleanExpiredItem() {
	// 设置一个定时触发器: 定时向 Ticker.C 中发送一个消息，即触发了一次
	timeTicker := time.NewTicker(mc.cleanExpiredItemTimeInterval)
	defer timeTicker.Stop()

	for {
		select {
		case <-timeTicker.C: // 每个周期做一个缓存清理
			// 遍历所有缓存的键值对，将有过期时间且过期的键删除
			for key, item := range mc.values {
				if item.expire != 0 && time.Now().After(item.expireTime) {
					mc.locker.Lock()
					mc.del(key)
					mc.locker.Unlock()
				}
			}
		}
	}
}

// 判断是否存在对应的 value
func (mc *memCache) get(key string) (*memCacheValue, bool) {
	val, ok := mc.values[key]
	return val, ok
}

// 删除: 当前内存大小更新、删除当前 key 值
func (mc *memCache) del(key string) bool {
	tmp, ok := mc.get(key)
	if ok && tmp != nil {
		mc.currMemorySize -= tmp.size
		delete(mc.values, key)
	}
	return true
}

// 添加: 当前内存大小更新、删除当前 key 值
func (mc *memCache) add(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currMemorySize += val.size
}

// 设置内存缓存的最大内存大小
func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	// fmt.Println("[DEBUG] set max memory size: ", mc.maxMemorySize, mc.maxMemorySizeStr)
	return true
}

// Set 将 value 写入缓存
func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	// fmt.Println("[DEBUG] set method")
	// map 非线程安全需要加锁访问
	mc.locker.Lock()
	defer mc.locker.Unlock()

	// 确定一个 value 值
	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		expire:     expire,
		size:       GetValSize(val),
	}

	// 为了简化代码复杂度, 这里用 “删除再添加” 来代替 “更新” 操作
	if _, ok := mc.get(key); ok { // 存在则删除
		mc.del(key)
	}

	// 新增后缓存是否超过最大内存: 超过则直接删除刚刚添加的这个 key, 并报 panic
	if mc.currMemorySize+v.size > mc.maxMemorySize {
		// 这里可以自己完善一下，通过一些内存淘汰策略来选择删除一些 key，来判断是否还会超过最大内存
		log.Println(fmt.Sprintf("[WARN] mem cache size is over max memory size: %d", mc.maxMemorySize))
		return false
	}
	mc.add(key, v)

	return true
}

// Get 根据 key 值获取 value
func (mc *memCache) Get(key string) (interface{}, bool) {
	mc.locker.RLock()
	defer mc.locker.RUnlock()

	// 拿到对应的值
	mcv, ok := mc.get(key)
	// 判断是否过期
	if ok {
		// 过期时间早于当前时间，删除
		if mcv.expire != 0 && mcv.expireTime.Before(time.Now()) {
			mc.del(key)
			return nil, false
		}

		return mcv.val, ok
	}
	return nil, false
}

// Del 删除 key 值
func (mc *memCache) Del(key string) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.del(key)
	return true
}

// Exists 判断 key 是否存在
func (mc *memCache) Exists(key string) bool {
	mc.locker.RLock()
	defer mc.locker.RUnlock()

	// TODO: 可以添加验证指定的 key 是否存在，以及是否过期
	_, ok := mc.values[key]
	return ok
}

// Flush 清空所有 key
func (mc *memCache) Flush() bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	// 直接将整个 map 置空，go 的垃圾回收机制会自行将没有使用的内存进行回收
	mc.values = make(map[string]*memCacheValue, 0)
	mc.currMemorySize = 0

	return true
}

// Keys 获取缓存中所有 key 的数量
func (mc *memCache) Keys() int64 {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	return int64(len(mc.values))
}

// cache/memCache.go
func NewMemCache() Cache {
	mc := &memCache{
		values:                       make(map[string]*memCacheValue),
		cleanExpiredItemTimeInterval: time.Second, // 定期清理缓存
	}

	// 轮询检查删除过期键
	go mc.cleanExpiredItem()
	return mc
}
