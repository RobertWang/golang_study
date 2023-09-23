package main

import (
	"fmt"
	"net"

	"github.com/c-robinson/iplib"
)

func main() {
	n := iplib.NewNet4(net.ParseIP("192.168.1.0"), 24)
	fmt.Println(n.Contains(net.ParseIP("192.168.1.3"))) // true
	fmt.Println(n.Contains(net.ParseIP("192.168.2.3"))) // false
}
