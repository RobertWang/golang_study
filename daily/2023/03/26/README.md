# Go 每日一题

> 今日（2023-03-26）的题目如下

下面的两个切片声明中有什么区别？哪个更可取？

```golang
A. var a []int
B. a := []int{}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：第一种切片声明不会分配内存，优先选择。

</div>
</details>
