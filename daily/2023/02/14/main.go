package main

import (
	"sync"
)

const N = 10

var wg = &sync.WaitGroup{}

func main() {
	for i := 0; i < N; i++ {
		// 改在这个位置
		wg.Add(1)

		go func(i int) {
			// 在这里会偶发引起下面的 panic 异常
			// wg.Add(1)
			println(i)
			defer wg.Done()
			// wg.Done()
		}(i)
	}
	// 可能偶发 panic : sync: WaitGroup is reused before previous Wait has returned
	wg.Wait()
}
