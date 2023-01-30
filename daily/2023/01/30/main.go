package main

// 28 楼回复

import (
	"fmt"
)

type DemoStruct struct {
	name string
}

func main() {
	var p = &DemoStruct{
		name: "tom",
	}
	fmt.Printf("A: %v\n", p.name)    // invalid operation: A (variable of type StructA) is not an interfacecompilerInvalidAssert
	fmt.Printf("C: %v\n", (*p).name) // invalid operation: A (variable of type StructA) is not an interfacecompilerInvalidAssert
}
