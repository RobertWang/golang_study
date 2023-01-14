package main

import (
	"fmt"
)

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	var a A = Work{3}
	// 类型断言
	// http://c.biancheng.net/view/4281.html
	s := a.(Work)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}
