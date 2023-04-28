# Go 每日一题

> 今日（2023-04-28）的题目如下

关于 switch 语句，下面说法正确的有?

- A. 条件表达式必须为常量或者整数；
- B. 单个 case 中，可以出现多个结果选项；
- C. 需要用 break 来明确退出一个 case；
- D. 只有在 case 中明确添加 fallthrough 关键字，才会继续执行紧跟的下一个 case；

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：BD。

---

### 9 楼

我觉得和其他语言最大的区别就是，case 里面可以放表达式，可以写出很简洁的代码。

```golang
package main

func main() {
    a := 1
    b := 0

    switch a {
    case b + 1: // 表达式
    case f(), b + 2: // 函数也行,多值都行
    }
}

func f() int {
    return 1
}
```

</div>
</details>
