package main

import (
	"fmt"
)

func main() {
	// var x string = nil
	// if x == nil {
	// 	x = "default"
	// }
	// go run main.go
	// # command-line-arguments
	// ./main.go:8:17: cannot use nil as string value in variable declaration
	// ./main.go:9:10: invalid operation: x == nil (mismatched types string and untyped nil)

	var x string
	var y string = ""
	fmt.Printf("x is %T, %v, %v\n", x, x, len(x))
	fmt.Printf("y is %T, %v, %v\n", y, y, len(y))

	if len(x) == 0 {
		x = "default"
	}
	fmt.Println(x)
}
