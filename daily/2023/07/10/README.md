# Go 每日一题

> 今日（2023-07-10）的题目如下

下面这段代码输出什么？

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
	var a A = Work{3}
	s := a.(Work)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}
```

- A. 13 23
- B. compilation error

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：A。

知识点：类型断言。这道题可以和第 15 天的第三题 和第 16 天的第三题结合起来看。

</div>
</details>
