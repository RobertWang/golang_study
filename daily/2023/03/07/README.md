# Go 每日一题

> 今日（2023-03-07）的题目如下

对 add() 函数调用正确的是？

```golang
func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}
```

- A. add(1, 2)
- B. add(1, 3, 7)
- C. add([]int{1, 2})
- D. add([]int{1, 3, 7}...)

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：ABD。

知识点：可变函数。

</div>
</details>
