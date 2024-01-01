package main

import (
	"fmt"
	"plugin"
)

func main() {
	fmt.Println("[so] Main() Invoked!")

	ptr, err := plugin.Open("funcs.so")
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := ptr.Lookup("Hello")
	print(hello)
	hello.(func(string) string)("test")
}
