package main

import (
	"fmt"
	"time"
)

var a = "0"

func main() {

	ch := make(chan string)

	go func() {
		i := 1

		for {

			if i%2 == 0 {
				a = "0"
			} else {
				a = "aa"
			}
			// 如果将sleep去掉，则会读不到脏数据，原因在于编译器做了优化
			time.Sleep(1 * time.Millisecond)
			i++
		}
	}()

	go func() {
		for {
			b := a
			if b != "0" && b != "aa" {
				ch <- b
			}
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Got strange string: ", <-ch)
	}

}

/*
golang之string类型变量操作的原子性
https://dashen.tech/2018/07/19/golang%E4%B9%8Bstring%E7%B1%BB%E5%9E%8B%E5%8F%98%E9%87%8F%E6%93%8D%E4%BD%9C%E7%9A%84%E5%8E%9F%E5%AD%90%E6%80%A7/

https://godbolt.org/
*/
