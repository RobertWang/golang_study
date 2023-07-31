# Go每日一题

> 今日（2022-12-07） 的题目如下

下面 A、B 两处应该填入什么代码，才能确保顺利打印出结果？

```golang
package main

import (
	"fmt"
)

type S struct {
	m string
}

func f() *S {
	return __  //A
}

func main() {
	p := __    //B
	fmt.Println(p.m) //print "foo"
}
```

自测答案

```
A: &S{"foo"}
B: &f()

./main.go:23:8: invalid operation: cannot take address of f() (value of type *S)
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：

A. `&S{"foo"}`
B. `*f()` 或者 `f()`

f() 函数返回参数是指针类型，所以可以用 & 取结构体的指针；
B 处，如果填 `*f()`，则 p 是 S 类型；如果填 `f()`，则 p 是 `*S` 类型，不过都可以使用 `p.m` 取得结构体的成员。

</div>
</details>