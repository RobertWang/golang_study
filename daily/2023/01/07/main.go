package main

import (
	"fmt"
)

type Integer int

// A test : Passed
func (a Integer) Add(b Integer) Integer {
	return a + b
}

// B test :
// ./main.go:24:26: cannot use b (variable of type Integer) as type *Integer in argument to i.(*Integer).Add
//
//	func (a Integer) Add(b *Integer) Integer {
//		return a + *b
//	}
//
// C test : Passed
//
//	func (a *Integer) Add(b Integer) Integer {
//		return *a + b
//	}
//
// D test :
// ./main.go:36:26: cannot use b (variable of type Integer) as type *Integer in argument to i.(*Integer).Add
//
//	func (a *Integer) Add(b *Integer) Integer {
//		return *a + *b
//	}

func main() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b)
	fmt.Println(sum)
}
