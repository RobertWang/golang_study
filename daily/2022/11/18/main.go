package main

import (
    "fmt"
)

func main() {
    println("2022-11-18 练习题")
    println("1: ------\n")
    a := []int{2: 1}
    fmt.Println(a)
    // [0 0 1]
    // println(a)
    // [3/3]0xc000044758

    println("2: ------\n")
    var x = []int{4: 44, 55, 66, 1: 77, 88}
    println(len(x), x[2])
    // 7 88
}

