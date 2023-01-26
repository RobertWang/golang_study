package main

import (
	"fmt"
)

func main() {
	fmt.Println("fix method 1:")
	list := new([]int)
	// current <= 通过 new([]int) 初始化的是个指针，而指针无法使用 append
	// list = append(list, 1)
	*list = append(*list, 1)
	fmt.Println(*list)

	fmt.Println("or use method 2:")
	list2 := make([]int, 0)
	list2 = append(list2, 1)
	fmt.Println(list2)
}
