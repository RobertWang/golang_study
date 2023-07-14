package main

import "fmt"

func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}

func main() {
	f()
	fmt.Println("M")
}
