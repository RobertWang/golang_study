# Go 每日一题

> 今日（2023-06-26）的题目如下

下面代码输出什么？

```golang
func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p :=1
	incr(&p)
	fmt.Println(p)
}
```

- A. 1
- B. 2
- C. 3

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：B。

知识点：指针，incr() 函数里的 p 是 *int 类型的指针，指向的是 main() 函数的变量 p 的地址。 第 2 行代码是将该地址的值执行一个自增操作，incr() 返回自增后的结果。

</div>
</details>
