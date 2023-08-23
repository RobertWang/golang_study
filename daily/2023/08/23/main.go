package main

import (
	"fmt"
	"sync"
	"time"
)

func demo1() {
	fmt.Println("demo 1: 常规遍历删除")
	fmt.Println("- 初始化")
	var m = make(map[string]int, 0)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	fmt.Printf("  初始化后的变量 m 为 %v\n", m)
	fmt.Printf("  map 长度为 %v\n", len(m))

	fmt.Println("- 删除操作...")
	for k := range m {
		key := k
		fmt.Println("  查看并删除 key : ", key, ", value ", m[key])
		delete(m, key)
	}
	fmt.Printf("  删除后的变量 m 为 %v\n", m)
	fmt.Printf("  map 长度为 %v\n", len(m))

	fmt.Println("")
}

func demo2() {
	fmt.Println("demo 2: 协程操作时删除")
	fmt.Println("- 初始化")
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4,
	}
	fmt.Printf("  初始化后的变量 m 为 %v\n", m)
	fmt.Printf("  map 长度为 %v\n", len(m))

	fmt.Println("- 删除操作...")
	for k := range m {
		key := k
		go func(delkey string) {
			fmt.Println("  正在删除的 key : ", delkey, ", value : ", m[delkey])
			delete(m, delkey)
		}(key)
	}
	fmt.Println("    延迟 3 秒...")
	time.Sleep(3 * time.Second)
	fmt.Printf("  删除后的变量 m 为 %v\n", m)
	fmt.Printf("  map 长度为 %v\n", len(m))
	fmt.Println("")
}

func demo3() {
	fmt.Println("demo 3: 重新定义变量")
	fmt.Println("- 初始化")
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4,
	}
	fmt.Printf("  初始化后的变量 m 为 %v\n", m)
	fmt.Printf("  map 长度为 %v\n", len(m))

	fmt.Println("- 删除操作...")
	m = make(map[string]int)
	fmt.Printf("  删除后的变量 m 为 %v\n", m)
	fmt.Printf("  map 长度为 %v\n", len(m))
	fmt.Println("")
}

func demo4() {
	fmt.Println("demo 4: 使用 sync.Map")
	fmt.Println("- 初始化")
	var m sync.Map
	m.Store("a", 1)
	m.Store("b", "test")
	m.Store("c", false)
	m.LoadOrStore("c", true)
	m.LoadOrStore("d", nil)
	fmt.Printf("  初始化后的变量 m 为 %v\n", m)
	// fmt.Printf("  map 长度为 %v\n", len(m))

	fmt.Println("- 删除操作...")
	walk := func(key, value interface{}) bool {
		fmt.Println("  正在删除的 key : ", key, ", value : ", value)
		m.Delete(key)
		return true
	}
	m.Range(walk)
	fmt.Printf("  删除后的变量 m 为 %v\n", m)
	// fmt.Printf("  map 长度为 %v\n", len(m))
	fmt.Println("")
}

func main() {
	// 01: 基本方式
	demo1()

	// 02: 协程访问
	demo2()

	// 03: 重新定义变量
	demo3()

	// 04: 使用 sync.Map
	demo4()
}

// 执行方式
// go run -race main.go
