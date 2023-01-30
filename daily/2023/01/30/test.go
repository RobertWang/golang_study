package main

// 28 楼回复

import (
	"fmt"
)

type StructA struct {
	P StructP
}

type StructP struct {
	Name string
}

func main() {
	var p = StructP{
		Name: "tom",
	}
	var A = StructA{
		P: p,
	}
	fmt.Printf("A: %v\n", A.(*P).Name) // invalid operation: A (variable of type StructA) is not an interfacecompilerInvalidAssert
}
