package main

import (
	"fmt"
)

// 这里的 func Name 需要大写，以符合 golang 的标准
func Hello(x string) string {
	fmt.Println("[so] Hello, Function!")
	return x
}

func hellox() {
	fmt.Println("[so] Hello, Function-X!")
}
