# Go每日一题

> 今日（2023-02-18）的题目如下

下面这段代码输出什么？

```golang
func main() {  
    a := 5
    b := 8.1
    fmt.Println(a + b)
}
```

- A. 13.1
- B. 13
- C. compilation error


<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：C。

a 的类型是 int，b 的类型是 float，两个不同类型的数值不能相加，编译报错。

---

### 1楼

Go语言的类型机制更加严格，没有隐式类型转换，所以不同类型的数据不能直接参与同一个运算。


</div>
</details>
