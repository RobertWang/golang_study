# Go每日一题

> 今日（2023-01-05） 的题目如下

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
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：C。

知识点：强制类型转化。

</div>
</details>