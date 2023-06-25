# Go 每日一题

> 今日（2023-06-25）的题目如下

下列选项正确的是？

```golang
func main() {
	str := "hello"
	str[0] = 'x'
	fmt.Println(str)
}
```

- A. hello
- B. xello
- C. compilation error

<details>
<summary>答案解析：</summary>
<div>

参考代码及解析：C。

知识点：Go 语言中的字符串是只读的。

</div>
</details>
