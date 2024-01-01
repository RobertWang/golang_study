// cache/cache.go
package cache

import (
	"time"
)

type Cache interface {
	// SetMaxMemory size : 1KB 100KB 1MB 2MB 1GB
	SetMaxMemory(size string) bool
	// Set 将 value 写入缓存
	Set(key string, val interface{}, expire time.Duration) bool
	// Get 根据 key 值获取 value
	Get(key string) (interface{}, bool)
	// Del 删除 key 值
	Del(key string) bool
	// Exists 判断 key 是否存在
	Exists(key string) bool
	// Flush 清空所有 key
	Flush() bool
	// Keys 获取缓存中所有 key 的数量
	Keys() int64
}
