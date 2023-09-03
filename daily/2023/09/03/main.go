package main

const s = "Go101.org"

// len(s) == 9
// 1 << 9 == 512
// 512 / 128 == 4

var v = "Go101.org"
var sb = [9]byte{'G', 'o', '1', '0', '1', '.', 'o', 'r', 'g'}

var a byte = 1 << len(s) / 128 // 是常量位移表达式 a 会变为 int 即 uint32 或 uint64 (0~4294967295/18446744073709551615) 所以不会溢出
var b byte = 1 << len(s[:]) / 128 // 不是常量位移表达式 b 仍为 byte 即 uint8 (0~255) 会溢出即为 0

var c byte = 1 << len(v) / 128
var d byte = 1 << len(v[:]) / 128

var e byte = 1 << len(sb) / 128
var f byte = 1 << len(sb[:]) / 128

func main() {
	println("拆解 a 的过程")
	println(s)
	println(len(s))
	println(1 << len(s))

	println("拆解 b 的过程")
	println(s[:])
	println(len(s[:]))
	println(1 << len(s[:]))

	println("原文的问题结果: a, b")
	println(a, b)
	println("扩展内容: c, d, e, f")
	println(c, d, e, f)
}

/*
Go 语言数据类型
https://www.runoob.com/go/go-data-types.html


*/