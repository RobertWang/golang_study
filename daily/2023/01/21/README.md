# Go每日一题

> 今日（2023-01-21） 的题目如下

以下代码输出什么？

```golang
package main

import (
	"fmt"
)

func main() {
	var a, b float64 = 1.0, 4.0
	fmt.Println(a | b)
}
```

- A：5
- B：+Inf
- C：panic
- D：不能编译


<details>
<summary>答案解析：</summary>
<div>

正确答案：D

`|` 操作是按位或操作符，它的操作数**只能是整数**，而上面这道题的操作数是 `float64`，因此编译不通过。

这是 Go 规范的内容 [https://docs.studygolang.com/ref/spec#Arithmetic_operators](https://docs.studygolang.com/ref/spec#Arithmetic_operators)：

```
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers

&    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&^   bit clear (AND NOT)    integers

<<   left shift             integer << integer >= 0
>>   right shift            integer >> integer >= 0
```

可以通过 play 在线编译看看：[https://play.studygolang.com/p/lLMbGE_ajrg](https://play.studygolang.com/p/lLMbGE_ajrg)

### 验证

VSCode 会提示以下错误

```
# command-line-arguments
./main.go:9:14: invalid operation: operator | not defined on a (variable of type float64)
```

</div>
</details>