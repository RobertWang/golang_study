# Go 每日一题

> 今日（2023-07-17）的题目如下

A、B、C、D 哪些选项有语法错误？

```golang
type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s) //A
	g(s) //B
	f(p) //C
	g(p) //D
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：BD。

函数参数为 interface{} 时可以接收任何类型的参数，包括用户自定义类型等，即使是接收指针类型也用 interface{}，而不是使用 \*interface{}。

>     永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。

</div>
</details>
