package main

import (
	"fmt"
)

func test1() {
	a := []int{2: 1}
	fmt.Println(a)
}

func test2() {
	var x = []int{4: 44, 55, 66, 1: 77, 88}
	fmt.Println(x)
	println(len(x), x[2])
}

/*
func test3() {
	// # command-line-arguments
	// ./main.go:21:38: duplicate index 4 in array or slice literal

	// 在 VSCode 编辑器中的信息提示 88 的位置红线
	// duplicate index 4 in array or slice literalcompilerDuplicateLitKey
	var x = []int{4: 44, 55, 66, 3: 77, 88}
	fmt.Println(x)
}
*/

func test4() {
	var x = []int{4: 44, 55, 66, 3: 77, 2: 88}
	x[2] = 2
	fmt.Println(x)
}

func main() {
	fmt.Println("测试一")
	test1()
	fmt.Println("")

	fmt.Println("测试二")
	test2()
	fmt.Println("")

	fmt.Println("测试四")
	test4()
	fmt.Println("")

}
