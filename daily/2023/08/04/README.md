# Go 每日一题

> 今日（2023-08-04）的题目如下

下面这段代码输出什么？

```golang
var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p)
}

func main() {
	p, err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}
```

- A. 5 5
- B. runtime error

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：B。知识点：变量作用域。问题出在操作符`:=`，对于使用`:=`定义的变量，如果新变量与同名已定义的变量不在同一个作用域中，那么 Go 会新定义这个变量。对于本例来说，main() 函数里的 p 是新定义的变量，会遮住全局变量 p，导致执行到 bar()时程序，全局变量 p 依然还是 nil，程序随即 Crash。

正确的做法是将 main() 函数修改为：

```golang
func main() {
	var err error
	p, err = foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}
```

这道题目引自 Tony Bai 老师的一篇文章，原文讲的很详细，推荐。

[https://tonybai.com/2015/01/13/a-hole-about-variable-scope-in-golang/](https://tonybai.com/2015/01/13/a-hole-about-variable-scope-in-golang/)

</div>
</details>
