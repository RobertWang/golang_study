# Go每日一题

> 今日（2022-12-01）的题目如下

通常，JS 面试，闭包应该是必考的题目。随着越来越多的语言对函数式范式的支持，闭包问题经常出现。在 Go 语言中也是如此。

这是 Go 语言爱好者周刊第 90 期的一道题目。以下代码输出什么？

```golang
package main

import "fmt"

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := app()
	fmt.Printf("a: %p\n", &a)
	b := app()
	fmt.Printf("b: %p\n", &b)
	a("go")
	fmt.Println(b("All"))
	fmt.Println(a("All"))
}
```

A：Hi All；B：Hi go All；C：Hi；D：go All


这道题目答对的人蛮多的：60%。不管你是答对还是答错，如果最后再加一行代码：fmt.Println(a("All"))，它输出什么？想看看你是不是蒙对了。（提示：你可以输出 t 的地址，看看是什么情况。）

<details>
<summary>自测结果</summary>
<div>

- a: 0xc000012028
- b: 0xc000012038
- Hi All
- Hi go All

</div>
<details>
