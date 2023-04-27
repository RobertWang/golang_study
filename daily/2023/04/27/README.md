# Go 每日一题

> 今日（2023-04-27）的题目如下

关于类型转化，下面选项正确的是？

```golang
A.
type MyInt int
var i int = 1
var j MyInt = i

B.
type MyInt int
var i int = 1
var j MyInt = (MyInt)i

C.
type MyInt int
var i int = 1
var j MyInt = MyInt(i)

D.
type MyInt int
var i int = 1
var j MyInt = i.(MyInt)
```

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：C。

知识点：强制类型转化。

---

### 17 楼

D 是类型断言

### 27 楼

本题的个人见解，请见下[链接](https://oyto.github.io/2023/04/27/Go%E6%AF%8F%E6%97%A5%E4%B8%80%E9%A2%98/%E7%B1%BB%E5%9E%8B%E8%BD%AC%E5%8C%96/)

### 28 楼

Golang type 类型别名和类型定义 [https://studygolang.com/articles/19144](https://studygolang.com/articles/19144)

</div>
</details>
