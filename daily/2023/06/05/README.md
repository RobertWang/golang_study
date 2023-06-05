# Go 每日一题

> 今日（2023-06-05）的题目如下

执行下面的代码会发生什么？

```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
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
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
```


<details>
<summary>答案解析：</summary>
<div>

记住 channel 的一些关键特性：

- 给一个 nil channel 发送数据，造成永远阻塞
- 从一个 nil channel 接收数据，造成永远阻塞
- 给一个已经关闭的 channel 发送数据，引起 panic
- 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
- 无缓冲的channel是同步的，而有缓冲的channel是非同步的

15字口诀：“**空读写阻塞，写关闭异常，读关闭空零**”，往已经关闭的 channel 写入数据会 panic。

本题中，因为 main 在开辟完两个 goroutine 之后，立刻关闭了 ch， 结果就是 panic：

>	panic: send on closed channel

---

### 9 楼

题解：过早的关闭 channel 导致了 panic

但是我有一个问题，将代码改正确，为什么 fmt.Println("close") 没有被打印呢？


```golang
ch := make(chan int, 1000)
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

fmt.Println("ok")
time.Sleep(time.Second * 1)
close(ch)
```

</div>
</details>
