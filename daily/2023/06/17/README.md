# Go 每日一题

> 今日（2023-06-17）的题目如下

下面这段代码输出什么？

```golang
func main() {  
    var i interface{}
    if i == nil {
        fmt.Println("nil")
        return
    }
    fmt.Println("not nil")
}
```

- A. nil
- B. not nil
- C. compilation error


<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

这是很常见的问题。

参考答案及解析：A。

当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil。

</div>
</details>
