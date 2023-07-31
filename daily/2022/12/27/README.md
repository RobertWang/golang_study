# Go每日一题

> 今日（2022-12-27） 的题目如下

下面这段代码输出什么？

```golang
func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(f(3))
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：7。

[Golang中的Defer必掌握的7知识点](https://studygolang.com/articles/27408)

</div>
</details>