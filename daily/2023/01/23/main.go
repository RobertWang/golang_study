package main

const s = "Go101.org"

// len(s) == 9
// 1 << 9 == 512
// 512 / 128 == 4

var a byte = 1 << len(s) / 128
var b byte = 1 << len(s[:]) / 128

func main() {
	print(len(s), "\n")
	print(len(s[:]), "\n")
	println(a, b)
}
