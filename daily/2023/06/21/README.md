# Go 每日一题

> 今日（2023-06-21）的题目如下

定义一个包内全局字符串变量，下面语法正确的是（多选）：

- A. var str string
- B. str := ""
- C. str = ""
- D. var str = ""

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：AD。

B 只支持局部变量声明；C 是赋值，str 必须在这之前已经声明。


---

### 35 楼

测试代码

```golang
package main

import "fmt"

//var str string
//var str = ""

//str =""
str:=""
func main() {
    fmt.Println("in main str:", str)
    callGlobalString()
}

func callGlobalString() {
    fmt.Println("in callGlobalString str ", str)
}
```

</div>
</details>
