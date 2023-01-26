package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	// s1 = append(s1, s2)
	// 向一个切片 s1 中追加另一个切片 s2 的数据时，需要对 s2 进行拆包
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
