package main

import (
	"encoding/hex"
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	s := "123"
	// print(&s)
	dump("s1", &s)
	s = "test"
	// print(&s)
	// print(s)
	dump("s2", &s)
}

func dump(name string, strPtr *string) {
	// 分隔符
	fmt.Println(strings.Repeat("-", 30))

	// 打印strPtr这个指针类型的变量的内存地址
	fmt.Printf("&%s:%p\n", name, strPtr)
	// 强制转换为16字节的byte数组 （*[0x10]byte)(unsafe.Pointer(strPtr)），是个指针类型。再用*取指针内容； 然后将数组转为切片，[:]
	fmt.Println(hex.Dump((*(*[0x10]byte)(unsafe.Pointer(strPtr)))[:]))

	// 获取上面这个指针所指向的内存地址
	p := *(*int)(unsafe.Pointer(strPtr)) // 当成int型(需要C经验更好理解)
	fmt.Printf("%s:0x%x\n", name, p)
	fmt.Println(hex.Dump((*(*[0x20]byte)(unsafe.Pointer(uintptr(p))))[:]))

}

/*
------------------------------
&s1:0xc000096220
00000000  2b 86 0a 01 00 00 00 00  03 00 00 00 00 00 00 00  |+...............|

s1:0x10a862b
00000000  31 32 33 31 32 35 36 32  35 3f 3f 3f 45 4f 46 48  |123125625???EOFH|
00000010  61 6e 4c 61 6f 4d 61 79  4d 72 6f 4e 61 4e 4e 6b  |anLaoMayMroNaNNk|

------------------------------
&s2:0xc000096220
00000000  7e 87 0a 01 00 00 00 00  04 00 00 00 00 00 00 00  |~...............|

s2:0x10a877e
00000000  74 65 73 74 74 72 75 65  75 69 6e 74 20 2e 2e 2e  |testtrueuint ...|
00000010  0a 20 4d 42 2c 20 20 61  6e 64 20 20 63 6e 74 3d  |. MB,  and  cnt=|

结构如下
s1 := StringHeader{
	Data: `指向"aa"的那个内存地址，记为x`,
	Len: 1,
}

s2 := StringHeader{
	Data: `指向"0"的那个内存地址，记为y`,
	Len: 2,
}


## 结论:

`string` 类型并不是并发安全的.
在并发场景下，string跟map一样，都是需要使用 atomic包/sync 包来保证读写的原子性。

*/
