# Go 每日一题

> 今日（2023-04-29）的题目如下

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

则 Add 函数定义正确的是：

```golang
A.
type Integer int
func (a Integer) Add(b Integer) Integer {
        return a + b
}

B.
type Integer int
func (a Integer) Add(b *Integer) Integer {
        return a + *b
}

C.
type Integer int
func (a *Integer) Add(b Integer) Integer {
        return *a + b
}

D.
type Integer int
func (a *Integer) Add(b *Integer) Integer {
        return *a + *b
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：AC。

知识点：类型断言、方法集。

---

### 2 楼

go 中有些的变量不可以寻址，指的是不能通过&获得其地址。

所以 func( `*A` ) 只能接收 `*A`, func( A ) 可以接收 A 或者 `*A` ,通过指针一定能得到变量的值 `*A` -> A

### 4 楼

mark：`func(A)`可以接收 A 和 A，`func(A)`只能 A，因为有些变量不可寻址`(&获取地址)`

### 13 楼

go 中有些的变量不可以寻址，指的是不能通过&获得其地址。

所以 func( A ) 只能接收 A, func( A ) 可以接收 A 或者 A ,通过指针一定能得到变量的值 A -> A

还比如 map 里面 的 value 也是不可寻地址的，因为 map 扩容后，value 地址就会改变

</div>
</details>
