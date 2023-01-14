# Go每日一题

> 今日（2022-12-25） 的题目如下

下面这段代码能否正常结束？

```golang
func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
}
```

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：不会出现死循环，能正常结束。

循环次数在循环开始前就已经确定，循环内改变切片的长度，不影响循环次数。

</div>
</details>