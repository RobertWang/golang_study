package main

import (
	"fmt"
)

func main() {
	var x string = nil
	if x == nil {
		x = "default"
	}
	fmt.Println(x)
}
