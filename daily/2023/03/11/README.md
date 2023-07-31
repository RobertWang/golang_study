# Go 每日一题

> 今日（2023-03-11）的题目如下

下面这段代码输出什么？

```golang
func main() {  
    i := 65
    fmt.Println(string(i))
}
```

- A. A
- B. 65
- C. compilation error

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：A。

UTF-8 编码中，十进制数字 65 对应的符号是 A。


---

### 12楼

直接string(int)会按utf8转换


### 14楼

int 自动转为 rune 类型了。


### 20楼

string中int转成byte类型


</div>
</details>
