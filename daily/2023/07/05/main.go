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
	c := Work{3}
	var a A = c
	var b B = c
	// 下面这两行会有报错
	fmt.Println(a.ShowB())
	fmt.Println(b.ShowA())
	/*
		fmt.Println(a.ShowA())
		fmt.Println(b.ShowB())
		fmt.Println(c.ShowA(), c.ShowB())
		a 有 ShowA()
		b 有 ShowB()
		c 有 ShowA() 和 ShowB()
	*/
}
