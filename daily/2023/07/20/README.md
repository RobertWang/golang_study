# Go 每日一题

> 今日（2023-07-20）的题目如下

下面的代码有几处语法问题，各是什么？

```golang
package main
import (
    "fmt"
)
func main() {
    var x string = nil
    if x == nil {
        x = "default"
    }
    fmt.Println(x)
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：2 处有语法问题。

golang 的字符串类型是不能赋值 nil 的，也不能跟 nil 比较。

</div>
</details>
