# Go每日一题

> 今日（2023-01-31）的题目如下

下面这段代码能否通过编译？如果通过，输出什么？

```golang
package main

import "fmt"

type MyInt1 int
type MyInt2 = int

func main() {
	var i int =0
	var i1 MyInt1 = i 
	var i2 MyInt2 = i
	fmt.Println(i1,i2)
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案：编译不通过，cannot use i (type int) as type MyInt1 in assignment

参考解析：这道题考的是类型别名与类型定义的区别。

第 5 行代码是基于类型 int 创建了新类型 MyInt1，第 6 行代码是创建了 int 的类型别名 MyInt2，注意类型别名的定义时 = 。所以，第 10 行代码相当于是将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译当然不通过；而 MyInt2 只是 int 的别名，本质上还是 int，可以赋值。

第 10 行代码的赋值可以使用强制类型转化 var i1 MyInt1 = MyInt1(i)。

### VSCode报异常

```golang
var i int
cannot use i (variable of type int) as MyInt1 value in variable declarationcompilerIncompatibleAssign
```



</div>
</details>