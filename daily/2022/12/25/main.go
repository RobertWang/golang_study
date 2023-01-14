package main

import (
	"fmt"
)

func main() {
	v := []int{1, 2, 3}
	fmt.Println(v)
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}
