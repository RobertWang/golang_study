# Go 每日一题

> 今日（2023-03-15）的题目如下

下面代码输出什么？

```golang
type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowB())
	fmt.Println(b.ShowA())
}
```

- A. 23 13
- B. compilation error


<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：B。

知识点：接口的静态类型。a、b 具有相同的动态类型和动态值，分别是结构体 work 和 {3}；a 的静态类型是 A，b 的静态类型是 B，接口 A 不包括方法 ShowB()，接口 B 也不包括方法 ShowA()，编译报错。

看下编译的错误：

```
a.ShowB undefined (type A has no field or method ShowB)
b.ShowA undefined (type B has no field or method ShowA)
```

</div>
</details>
