package main

import "fmt"

type Math struct {
	x, y int
}

var m = map[string]Math{
	/*
		redundant type from array, slice, or map composite
		literal simplifycompositelit
	*/
	// "foo": Math{2, 3},
	"foo": {2, 3},
}

func main() {
	// m["foo"].x = 4
	fmt.Println(m["foo"].x)
}
