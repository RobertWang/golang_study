package main

import "fmt"

func (i int) PrintInt() {
	fmt.Println(i)
}

func main() {
	var i int = 1
	i.PrintInt()
}
