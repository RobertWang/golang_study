package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func main() {
	//定义map
	m := make(map[string]*student)

	// 定义student数组
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	// // 将数组依次添加到map中
	// for _, stu := range stus {
	// 	m[stu.Name] = &stu
	// }

	// 正确写法
	// 遍历结构体数组，依次赋值给map
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}

	// 或用中间变量方式，保证正确性，但会比较消耗性能
	// for _, stu := range stus {
	// 	tmpStu := stu
	// 	m[stu.Name] = &tmpStu
	// }

	// 打印map
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}
