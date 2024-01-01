// main.go

package main

import (
	cacheServer "main/server"	
)

func main() {

	cache := cacheServer.NewMemCache()
	cache.SetMaxMemory("10MB")
	cache.Set("int", 1)
	cache.Set("bool", false)
	cache.Set("data", map[string]interface{}{"a": 1})
	cache.Get("int")
	cache.Del("int")
	cache.Flush()
	cache.Keys()
}
