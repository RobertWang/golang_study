# Go 每日一题

> 今日（2023-06-08的题目如下

下面这段代码输出什么？

```golang
type person struct {  
    name string
}

func main() {  
    var m map[person]int
    p := person{"mike"}
    fmt.Println(m[p])
}
```

- A. 0
- B. 1
- C. Compilation error

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：A。

m 是一个 map，值是 nil。从 nil map 中取值不会报错，而是返回相应的零值，这里值是 int 类型，因此返回 0。

</div>
</details>
