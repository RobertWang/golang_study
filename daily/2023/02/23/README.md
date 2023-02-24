# Go 每日一题

> 今日（2023-02-23）的题目如下

下面这段代码输出什么？

```golang
func main() {
	a := [2]int{5, 6}
	b := [3]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
```

- A. compilation error
- B. equal
- C. not equal

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：A。

Go 中的数组是值类型，可比较，另外一方面，数组的长度也是数组类型的组成部分，所以 a 和 b 是不同的类型，是不能比较的，所以编译错误。

</div>
</details>
