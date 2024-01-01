package cache

import (
	"testing"
	"time"
)

func TestCacheOP(t *testing.T) {
	testData := []struct {
		key    string
		val    interface{}
		expire time.Duration
	}{
		{"integer", 678, time.Second * 10},
		{"bool1", false, time.Second * 11},
		{"bool2", true, time.Second * 12},
		{"dict", map[string]interface{}{"a": 3, "b": false}, time.Second * 13},
		{"strr", "info", time.Second * 14},
		{"chinese", "圆周率 3.14", time.Second * 15},
	}

	c := NewMemCache()
	c.SetMaxMemory("10MB")
	// 测试 set 和 get
	for _, item := range testData {
		c.Set(item.key, item.val, item.expire)
		val, ok := c.Get(item.key)
		if !ok {
			t.Error("缓存取值失败")
		}

		if item.key != "dict" && val != item.val {
			t.Error("缓存取值数据与预期不一致")
		}
		_, ok1 := val.(map[string]interface{})
		if item.key == "dict" && !ok1 {
			t.Error("缓存取值数据与预期不一致")
		}
	}

	// 测试 Keys()
	if int64(len(testData)) != c.Keys() {
		t.Error("缓存数量不一致")
	}

	// 测试 Del()
	c.Del(testData[0].key)
	c.Del(testData[1].key)
	if int64(len(testData)) != c.Keys()+2 {
		t.Error("缓存数量不一致")
	}

	// 测试过期时间
	time.Sleep(time.Second * 16)
	if c.Keys() != 0 {
		t.Error("缓存清空失败")
	}

}
