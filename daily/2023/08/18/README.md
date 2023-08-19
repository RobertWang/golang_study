# Go 每日一题

> 今日（2023-08-18）的题目如下

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

所以 `func( *A )` 只能接收 `*A`, `func( A )` 可以接收 `A` 或者 `*A` ,通过指针一定能得到变量的值 `*A` -> `A`

### 4 楼

mark：func(A）可以接收 A 和 A，func(A)只能 A，因为有些变量不可寻址（&获取地址）

### 13 楼

go 中有些的变量不可以寻址，指的是不能通过&获得其地址。

所以 `func( *A )` 只能接收 `*A`, `func( A )` 可以接收 `A` 或者 `*A` ,通过指针一定能得到变量的值 `*A` -> `A`

还比如 map 里面 的 value 也是不可寻地址的，因为 map 扩容后，value 地址就会改变

### 16 楼

> As with selectors, a reference to a non-interface method with a value receiver using a pointer will automatically dereference that pointer: pt.Mv is equivalent to (\*pt).Mv.
>
> As with method calls, a reference to a non-interface method with a pointer receiver using an addressable value will automatically take the address of that value: t.Mp is equivalent to (&t).Mp.

仅限于 selector 和 method(receiver)，普通参数可没有这待遇。而且这里面还涉及 method set 和 wrapper method，想扯的话还是挺多的。

至于你说的因为 `map 扩容后，value 地址就会改变`并不是理由，至少不够充分。

有没有想过栈也会扩容为什么就可以寻址了呢。

你不能说因为栈上 `adjustpointer` 比较容易所以可以寻址，map 比较难那就算了吧。

既然栈可以寻址，函数返回值难道不是在栈上（go1.17 之后有点模糊，但之前可是很确定的），怎么就不行了呢？

</div>
</details>
