# Go 每日一题

> 今日（2023-05-12）的题目如下

下面这段代码能否通过编译，如果可以，输出什么？

```golang
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
}
```

<details>
<summary>答案解析：</summary>
<div>

不能通过编译。append() 的第二个参数不能直接使用 slice，需使用 … 操作符，将一个切片追加到另一个切片上：append(s1,s2…)。或者直接跟上元素，形如：append(s1,1,2,3)。

---

### 4 楼

不行，s2需要 unpack： `s1 = append(s1,s2...)`

### 20 楼

append() 的第二个参数不能直接使用 slice，需使用 … 操作符

### 28 楼

在 Golang 中，参数 ...（称为 ellipsis 或 variadic）表示可以接收任意数量的参数，这些参数将被转换为一个切片，以在函数中使用。

```golang
func myFunc(args ...int) {
    for _, arg := range args {
        fmt.Println(arg)
    }
}

myFunc(1, 2, 3) // 输出: 1 2 3
```

在 Golang 中，... 用于将 slice 展开为可变参数列表，使其能够传递给期望接收可变参数的函数。在这种情况下，... 是一个语法糖，它允许我们将 slice 的每个元素视为一个单独的参数。因此，当我们调用一个期望接收可变参数的函数时，可以传递一个 slice，并将其展开为一个参数列表。

如果在 slice 后面加上 ellipsis，例如 slice1.append(slice1, slice2...)，实际上是将 slice2 展开为一个参数列表，并将其作为参数传递给 append 函数。这将导致 slice2 中的每个元素都被添加到 slice1 中。

需要注意的是，... 语法只能用于最后一个参数，用于将 slice 展开为可变参数列表。如果在其他位置使用 ...，则会导致编译错误。

</div>
</details>
