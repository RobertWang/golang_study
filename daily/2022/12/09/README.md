# Go每日一题

> 今日（2022-12-09） 的题目如下

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

运行结果:

```
./main.go:13:17: cannot use nil as string value in variable declaration
./main.go:14:10: invalid operation: x == nil (mismatched types string and untyped nil)
```


<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：2 处有语法问题。

golang 的字符串类型是不能赋值 nil 的，也不能跟 nil 比较。

</div>
</details>
