package main

import (
	"fmt"
	"time"
)

func main() {

	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}
