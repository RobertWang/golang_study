# Go 每日一题

> 今日（2023-07-31）的题目如下

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

>     与 [2022-12-20](../../../2022/12/20/README.md) 及 [2023-04-10](../../04/10/README.md) 相同

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：South。知识点：iota 的用法、类型的 `String()` 方法。

根据 iota 的用法推断出 South 的值是 2；另外，如果类型定义了 `String()` 方法，当使用 `fmt.Printf()`、`fmt.Print()` 和 `fmt.Println()` 会自动使用 `String()` 方法，实现字符串的打印。

Reference:

- [https://wiki.jikexueyuan.com/project/the-way-to-go/10.7.html](https://wiki.jikexueyuan.com/project/the-way-to-go/10.7.html)
- [https://www.sunansheng.com/archives/24.html](https://www.sunansheng.com/archives/24.html)
- [https://yourbasic.org/golang/iota/](https://yourbasic.org/golang/iota/)

</div>
</details>
