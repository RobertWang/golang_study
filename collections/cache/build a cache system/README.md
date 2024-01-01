# 用 golang 实现一个简单的缓存

## 面试题内容

- 支持设置过期时间，精度到秒
- 支持设置最大内存，当内存超出时做出合理的处理
- 支持并发安全
- 按照以下接口要求实现

```golang
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
```

### 使用示例

```golang
cache := NewMemCache()
cache.SetMaxMemory("100MB")
cache.Set("int", 1)
cache.Set("bool", false)
cache.Set("data", map[string]interface{}{"a": 1})
cache.Get("int")
cache.Del("int")
cache.Flush()
cache.Keys()
```

## 参考

- [Go Test 单元测试简明教程](https://geektutu.com/post/quick-go-test.html)
- [手把手教你用 Go 语言实现缓存系统](https://www.yinkai.cc/post/a78e65ab014d188b51f57487f4e5c45c)
