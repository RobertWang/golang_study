package main

import (
	"fmt"
)

func change(s ...int) {
	s = append(s, 3)
}

func main() {
	// make => make(type, len, cap)
	// 参考 : [Golang 内置函数：make()](https://blog.csdn.net/yilovexing/article/details/121172745)
	// Slice make	=> make(slice type, length, capacity)
	// Map make		=> make(map type, length, capacity)
	// Channel make	=> make(channel type, buffer size, buffer capacity)
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)
}
