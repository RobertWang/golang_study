# Go 每日一题

> 今日（2023-02-27）的题目如下

下面这段代码输出什么？

```golang
func main() {  
    i := -5
    j := +5
    fmt.Printf("%+d %+d", i, j)
}
```

- A. -5 +5
- B. +5 +5
- C. 0 0

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：A。

%d表示输出十进制数字，+表示输出数值的符号。这里不表示取反。

</div>
</details>
