package main

import (
	"fmt"
)

/*
make(T[], len, cap) 创建切片，其中：

- T：切片的元素类型
- len：切片的实际长度
- cap：切片的最大容量

注意：当cap不传值的话，默认 len=cap。

举例：s := make(int[], 5)。此时 s 的值为 [0 0 0 0 0]
*/

func demo1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func demo2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}

func demo3() {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func demo4() {
	s := make([]int, 5, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}


func main() {
	demo1()
	demo2()
	demo3()
	demo4()
}
