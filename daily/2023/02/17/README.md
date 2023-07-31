# Go每日一题

> 今日（2023-02-17）的题目如下

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

参考答案及解析：A.18。

知识点：可变参数函数。

### 1楼

引用类型传参

### 14楼

可变参数是切片，切片是引用，所以func内赋值会带出来。

### 15楼

可变参数函数。可变参数是切片，切片是引用


</div>
</details>
