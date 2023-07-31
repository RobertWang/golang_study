# Go 每日一题

> 今日（2023-04-20）的题目如下

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

```
[1 2 0 0 0]
[1 2 3 0 0]
```

知识点：可变函数、append()操作。

Go 提供的语法糖`...`，可以将 slice 传进可变函数，不会创建新的切片。第一次调用 change() 时，append() 操作使切片底层数组发生了扩容，原 slice 的底层数组不会改变； 第二次调用 change() 函数时，使用了操作符`[i,j]`获得一个新的切片，假定为 slice1，
它的底层数组和原切片底层数组是重合的，不过 slice1 的长度、容量分别是 2、5，所以在 change() 函数中对 slice1 底层数组的修改会影响到原切片。

---

### 8 楼

切片作为参数进行传递，如果函数内放生扩容，那么会产生新的切片。已经不再是原来的那个切片

### 9 楼

> 8 楼

所以第二种也是如此，只不过他也改变了原生切片的值

### 12 楼

理解了，底层数组是底层数组，切片只是个片段

### 22 楼

```golang
func main() {

    s := make([]int, 4, 5)
    s[0] = 1
    s[1] = 2

    change(s...)
    fmt.Println(s)
    fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
    fmt.Println("-------------------------")
    change(s[0:2]...)
    fmt.Println(s)
    fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)

}


func change(s ...int) {
    fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
    s = append(s, 3)
    fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
}
```

输出：

```
len: 4  cap: 5 pointer: 0xc000012390
len: 5  cap: 5 pointer: 0xc000012390
[1 2 0 0]
len: 4  cap: 5 pointer: 0xc000012390
-------------------------
len: 2  cap: 5 pointer: 0xc000012390
len: 3  cap: 5 pointer: 0xc000012390
[1 2 3 0]
len: 4  cap: 5 pointer: 0xc000012390
```

?!?!

### 24 楼

func change(s ...int) { s = append(s,3) }

func main() { slice := make([]int,5,5) slice[0] = 1 slice[1] = 2 change(slice...) fmt.Println(slice) change(slice[0:2]...) fmt.Println(slice) } 切片底层数组 第一次调用 change() 时，append() 操作使切片底层数组发生了扩容，原 slice 的底层数组不会改变； 第二次调用 change() 函数时，使用了操作符[i,j]获得一个新的切片，假定为 slice1， 它的底层数组和原切片底层数组是重合的，不过 slice1 的长度、容量分别是 2、5，所以在 change() 函数中对 slice1 底层数组的修改会影响到原切片。

### 30 楼

> 22 楼

```golang
func main() {

    s := make([]int, 4, 5)
    s[0] = 1
    s[1] = 2

    change(s...)
    fmt.Println(s)
    fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), &s)
    fmt.Println("-------------------------")
    change(s[0:2]...)
    fmt.Println(s)
    fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), &s)

}

func change(s ...int) {
    fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), &s)
    fmt.Println(s)
    s = append(s, 3)
    fmt.Println(s)
    fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), &s)
}
```

输出

```
len: 4  cap: 5 pointer: 0xc000010048
[1 2 0 0]
[1 2 0 0 3]
len: 5  cap: 5 pointer: 0xc000010048
[1 2 0 0]
len: 4  cap: 5 pointer: 0xc000010030
-------------------------
len: 2  cap: 5 pointer: 0xc0000100a8
[1 2]
[1 2 3]
len: 3  cap: 5 pointer: 0xc0000100a8
[1 2 3 0]
len: 4  cap: 5 pointer: 0xc000010030
```

参考地址：[https://segmentfault.com/a/1190000042430248](https://segmentfault.com/a/1190000042430248)

</div>
</details>
