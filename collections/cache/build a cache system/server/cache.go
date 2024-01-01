// cache/cache.go
package server

import (
	"main/cache"
	"time"
)

// 代理层/适配层
type cacheServer struct {
	memCache cache.Cache
}

func NewMemCache() *cacheServer {
	return &cacheServer{
		memCache: cache.NewMemCache(),
	}
}

func (cs *cacheServer) SetMaxMemory(size string) bool {
	return cs.memCache.SetMaxMemory(size)
}

func (cs *cacheServer) Set(key string, val interface{}, expire ...time.Duration ) bool {
	// 过期时间默认值为 0 称
	expireTs := time.Second * 0
	if len(expire) > 0 {
		expireTs = expire[0]
	}

	return cs.memCache.Set(key, val, expireTs)
}

func (cs *cacheServer) Get(key string) (interface{}, bool)  {
	return cs.memCache.Get(key)
}


// Del 删除 key 值
func (cs *cacheServer) Del(key string) bool {
	return cs.memCache.Del(key)
}

// Exists 判断 key 是否存在
func (cs *cacheServer) Exists(key string) bool {
	return cs.memCache.Exists(key)
}

// Flush 清空所有 key
func (cs *cacheServer) Flush() bool {
	return cs.memCache.Flush()
}

// Keys 获取缓存中所有 key 的数量
func (cs *cacheServer) Keys() int64 {
	return cs.memCache.Keys()
}
