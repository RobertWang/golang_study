package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	var m map[person]int
	p := person{"mike"}
	// 如果需要指定 m 的字典值，需要使用下面的两行
	m = make(map[person]int)
	m[p] = 5

	fmt.Println(m[p])

	// 查看当前 person 结构体内容
	fmt.Println(p)

	// 用下面的方法也可以获取到指定的字典值
	// 1. 直接
	fmt.Println(m[person{"mike"}])
	// 或 2. 同值变量
	t1 := person{"mike"}
	fmt.Println(m[t1])
	// 或 2. 同值变量
	t2 := person{name: "mike"}
	fmt.Println(m[t2])

	// 查看所有字典值
	fmt.Println(m)
}

// 上面代码运行结果如下 :
// 5
// {mike}
// 5
// 5
// 5
// map[{mike}:5]
