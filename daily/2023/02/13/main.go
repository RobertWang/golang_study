package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	// 在这里添加延迟关闭 channel 的操作，取代下面的 close(ch)，避免 panic: "send on closed channel"
	defer close(ch)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()

	// 下面这行将报 panic 异常
	// "send on closed channel"
	// close(ch)

	// 也可以在这里改成下面的方式
	// defer close(ch)

	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
