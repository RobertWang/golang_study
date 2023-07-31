# Go 每日一题

> 今日（2023-06-23）的题目如下

下面这段代码输出什么?

```golang
func hello(i int) {  
    fmt.Println(i)
}
func main() {  
    i := 5
    defer hello(i)
    i = i + 10
}
```


<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：5。

这个例子中，hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5.

---

### 01 楼

Mark 本来以为会在执行defer的时候获取i的地址，在最后执行hello的时候才复制i的值。但看了解析原来是在defer的时候就会保存一份i的副本。 个人感觉从逻辑上应该是输出15更合适一些。不知道是不是golang做了简单处理，在defer语句的时候就创建了副本。

### 07 楼

golang 在 `defer` 时候就创建了副本, 所以是 5;

如果在 `defer` 函数内部直接引用, 而不是通过参数的形式传递,则结果是 15


</div>
</details>
