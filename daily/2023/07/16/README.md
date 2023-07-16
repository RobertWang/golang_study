# Go 每日一题

> 今日（2023-07-16）的题目如下

下面的两个切片声明中有什么区别？哪个更可取？

```golang
A. var a []int
B. a := []int{}
```

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：第一种切片声明不会分配内存，优先选择。

---

### 6 楼

```golang
func Test02_00(t *testing.T) {
	var a [] int
	b := []int{}

	fmt.Println("a size: %v\n", unsafe.Sizeof(a))
	fmt.Println("a len: %v\n", len(a))
	fmt.Println("a size: %v\n", cap(a))
	fmt.Println("b size: %v\n", unsafe.Sizeof(b))
	fmt.Println("b len: %v\n", len(b))
	fmt.Println("b size: %v\n", cap(b))
}
/*
go test -timeout 30s -run ...

== RUN Test02_00
a size: 24
a len: 0
a cap: 0
b size: 24
b len: 0
b cap: 0
--- PASS: Test02_00 (0.00s)
PASS
ok		exam/other/02	0.002s
*/
```

### 8 楼

nil slice 和 empty slice 的区别具体是什么呢 nil slice 中底层的 pointer 还没分配地址，但是 empty slice 已经分配地址了？空数组可以分配地址吗？

### 13 楼

A 只是分配了指针；B 进一步分配了内存。B 更可取。

### 23 楼

对于数组我判断为空只用 `len(xx) == 0`，不用管是不是 nil 还是空切片。

</div>
</details>
