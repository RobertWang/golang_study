package main

import (
	"fmt"
)

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}

/*
https://studygolang.com/interview/question
下面代码输出什么？
A. 1 1
B. 0 1
C. 1 0
D. 0 0
*/

/*
参考答案及解析：B。

知识点：defer、返回值。注意一下，increaseA() 的返回参数是匿名，increaseB() 是具名。关于 defer 与返回值的知识点，后面我会写篇文章详细分析，到时候可以看下文章的讲解。
*/
