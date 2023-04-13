package main

import "fmt"

type Direction int

const (
	North Direction = iota // 0
	East                   // 1
	South                  //2
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func main() {
	fmt.Println(South)
}
