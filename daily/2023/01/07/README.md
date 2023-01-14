# Go每日一题

> 今日（2023-01-07） 的题目如下

如果 Add() 函数的调用代码为：

```golang
func main() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b)
	fmt.Println(sum)
}
```

则Add函数定义正确的是：

A.
```golang
type Integer int
func (a Integer) Add(b Integer) Integer {
    return a + b
}
```

B.
```golang
type Integer int
func (a Integer) Add(b *Integer) Integer {
    return a + *b
}
```

C.
```golang
type Integer int
func (a *Integer) Add(b Integer) Integer {
    return *a + b
}
```

D.
```golang
type Integer int
func (a *Integer) Add(b *Integer) Integer {
    return *a + *b
}
```

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：AC。

知识点：类型断言、方法集。

> go 中有些的变量不可以寻址，指的是不能通过&获得其地址。
>
> 所以 `func(*A)` 只能接收 `*A`, `func(A)` 可以接收 `A` 或者 `*A` ,通过指针一定能得到变量的值 `*A` -> `A`

> `func(A)` 可以接收 `*A` 和 `A`，`func(*A)` 只能接收 `A`，因为有些变量不可寻址（&获取地址）

</div>
</details>