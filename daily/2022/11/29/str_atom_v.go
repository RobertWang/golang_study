package main

import (
	"github.com/davecgh/go-spew/spew"
	"unsafe"
)

func main() {

	var str string = "0"
	p := (*struct {
		str uintptr
		len int
	})(unsafe.Pointer(&str))

	spew.Dump(p)

}
