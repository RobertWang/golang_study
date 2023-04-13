# Go 每日一题

> 今日（2023-04-10）的题目如下

下面这段代码输出什么？

```golang
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func main() {
	fmt.Println(South)
}
```

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：

South 根据 iota 的用法推断 South 的值是 2;另外，如果定义了 String()方法,当使用 fmt.Printf()、fmt.Print() 和 fmt.Println()会自动使用 String()方法, 实现字符串打印

所以，最终输出

```
South
```

</div>
</details>
