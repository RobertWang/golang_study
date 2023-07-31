# Go 每日一题

> 今日（2023-06-09）的题目如下

下面这段代码输出什么？

```golang
func hello(num ...int) {  
    num[0] = 18
}

func main() {  
    i := []int{5, 6, 7}
    hello(i...)
    fmt.Println(i[0])
}
```

- A. 18
- B. 5
- C. Compilation error

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：A.18。知识点：可变参数函数。

</div>
</details>
