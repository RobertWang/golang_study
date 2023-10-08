# 素数生成

```golang
// Copyright 2009 The Go Authors.All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package main
package main

import "fmt"

// Send the sequence 2,3,4,...to channel 'ch'.
func generate(ch chan int) {

	for i := 2; ; i++ {

		ch <- i //Send 'i'to channel 'ch'.

	}
}

// Copy the values from channel 'in'to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {

	for {
		i := <-in // Receive value of new variable 'i'from 'in'.
		if i%prime != 0 {
			out <- i //Send 'i'to channel 'out'.
		}
	}
}

// The prime sieve:Daisy-chain filter processes together.
func main() {

	ch := make(chan int) //Create a new channel.
	go generate(ch)

	//start generate()as a goroutine.

	for {

		prime := <-ch

		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
```

generate 用于产生序列 filter 用于过滤这个数是否可以被某个素数整除

main 中的 for 循环 会根据每一个产出的素数 新建一个协程过滤，且这个协程作为 过滤链的 最后 1 步

大致流程是

- 收到数据 2，创建协程 过滤能被 2 整除的数，
- 使用 ch（即生成数据的 channel）作为数据输入，
- ch1 作为数据输出 协程收到数据 3，
- 不能被 2 整除，输出 3，
- prime 收到 3，
- 创建协程 过滤能被 3 整除的数，并以协程 2 的输出作为输入，
- 输出 channel 为 ch，
- 即 prime 获取的值改为了 协程 3 的输出
