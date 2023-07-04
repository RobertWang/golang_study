package main

import (
	"fmt"
)

func main() {
	// var m map[string]int //A
	var m = make(map[string]int) //A
	m["a"] = 1
	// if v := m["b"]; v != nil { //B
	if v, ok := m["b"]; ok { //B
		fmt.Println(v)
	}
	/*
		v, ok
		0 false
	*/
}
