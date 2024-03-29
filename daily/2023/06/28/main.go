package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	// origin
	// go fmt.Println(<-ch1)

	// fixed
	go func() {
		fmt.Println(<-ch1)
	}()

	ch1 <- 5
	time.Sleep(1 * time.Second)
}
