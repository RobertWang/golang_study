package main

import (
	"fmt"
	"net/netip"
)

func main() {
	p, err := netip.ParsePrefix(`10.187.102.0/24`)
	if err != nil {
		panic(err)
	}
	a, err := netip.ParseAddr(`10.187.102.8`)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Contains(a))
}
