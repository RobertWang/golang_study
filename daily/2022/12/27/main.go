package main

import (
	"fmt"
)

func f(n int) (r int) {
	fmt.Println("n = ", n)
	defer func() {
		fmt.Println("defer func :: 1 : n = ", n, ", r=", r)
		r += n
		fmt.Println("defer func :: 2 : n = ", n, ", r=", r)
		recover()
		fmt.Println("defer func :: 3 : n = ", n, ", r=", r)
	}()

	var f func()
	// fmt.Println("var func f : 1 = ", f)
	defer f()
	f = func() {
		fmt.Println("defer func f :: 1 : n = ", n, ", r=", r)
		r += 2
		fmt.Println("defer func f :: 2 : n = ", n, ", r=", r)
	}
	// fmt.Println("var func f : 2 = ", f)
	// defer f()

	fmt.Println("func f :: n = ", n, ", r=", r)
	return n + 1
}

func main() {
	fmt.Println(f(3))
}
