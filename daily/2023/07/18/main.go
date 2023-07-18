package main

import (
	"fmt"
)

type S struct {
	m string
}

func f() *S {
	return nil //A
}

func main() {
	p := &S{m: "foo"} //B
	fmt.Println("p type is %T", p)
	fmt.Println(p.m) //print "foo"
}

// A 处 : &S{"foo"}
// B 处 :
// 1. f()  --> p 的类型为 *S 即 &{foo}
// 2. *f() --> p 的类型为 S 即 {foo}
// 另类写法，绕过预定义的 f()
// A 处 : nil
// B 处 :
// 1. S{m: "foo"}  --> p 的类型为 S 即 {foo}
// 2. &S{m: "foo"} --> p 的类型为 *S 即 &{foo}
