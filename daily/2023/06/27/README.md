# Go 每日一题

> 今日（2023-06-27）的题目如下

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


---

### 4 楼

ABD。因为它要求传递若干 int，A、B 显然正确；C 是数组或切片，不对；D 把切片展开，其实就是若干 int，也对



</div>
</details>
