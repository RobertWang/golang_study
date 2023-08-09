# Go 每日一题

> 今日（2023-08-09）的题目如下

下面这段代码输出什么？

```golang
func change(s ...int) {
	s = append(s,3)
}

func main() {
	slice := make([]int,5,5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：

```golang
[1 2 0 0 0]
[1 2 3 0 0]
```

知识点：可变函数、append()操作。

Go 提供的语法糖`...`，可以将 slice 传进可变函数，不会创建新的切片。第一次调用 change() 时，append() 操作使切片底层数组发生了扩容，原 slice 的底层数组不会改变； 第二次调用change() 函数时，使用了操作符[i,j]获得一个新的切片，假定为 slice1，
它的底层数组和原切片底层数组是重合的，不过 slice1 的长度、容量分别是 2、5，所以在 change() 函数中对 slice1 底层数组的修改会影响到原切片。

</div>
</details>
