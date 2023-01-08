package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	name string
}

func (this *Test) Point() {
	fmt.Println(this.name)
}

func main() {
	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}

	fmt.Println("Origin")
	for _, t := range ts {
		tv := t
		fmt.Println(reflect.TypeOf(tv))
		fmt.Printf("Point Address : %p\n", &tv)
		defer tv.Point()
	}

	fmt.Println("test by len for index asc")
	for i := len(ts); i > 0; i-- {
		tv := ts[i-1]
		fmt.Println("index is : ", i)
		fmt.Println(reflect.TypeOf(tv))
		fmt.Printf("Point Address : %p\n", &tv)
		fmt.Printf("this.Name is : %s\n", tv.name)
		defer tv.Point()
	}
}
