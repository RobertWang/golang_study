# Go每日一题

> 今日（2023-01-17） 的题目如下

init() 函数是什么时候执行的？

<details>
<summary>答案解析：</summary>
<div>

init() 函数是 Go 程序初始化的一部分。Go 程序初始化先于 main 函数，由 runtime 初始化每个导入的包，初始化顺序不是按照从上到下的导入顺序，而是按照解析的依赖关系，没有依赖的包最先初始化。

每个包首先初始化包作用域的常量和变量（常量优先于变量），然后执行包的 init() 函数。同一个包，甚至是同一个源文件可以有多个 init() 函数。init() 函数没有入参和返回值，不能被其他函数调用，同一个包内多个 init() 函数的执行顺序不作保证。

一句话总结： import –> const –> var –> init() –> main()

示例：

```golang
package main

import "fmt"

func init()  {
	fmt.Println("init1:", a)
}

func init()  {
	fmt.Println("init2:", a)
}

var a = 10
const b = 100

func main() {
	fmt.Println("main:", a)
}
// 执行结果
// init1: 10
// init2: 10
// main: 10
```

答案解析来源：[init() 函数是什么时候执行的？](https://geektutu.com/post/qa-golang-2.html#Q1-init-%E5%87%BD%E6%95%B0%E6%98%AF%E4%BB%80%E4%B9%88%E6%97%B6%E5%80%99%E6%89%A7%E8%A1%8C%E7%9A%84%EF%BC%9F)



### 23楼

init() 函数没有入参和返回值，不能被其他函数调用，同一个包内多个 init() 函数的执行顺序不作保证。

一句话总结： import –> const –> var –> init() –> main()

</div>
</details>