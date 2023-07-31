# Go 每日一题

> 今日（2023-05-09）的题目如下

init() 函数是什么时候执行的？

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

`init()` 函数是 Go 程序初始化的一部分。Go 程序初始化先于 main 函数，由 runtime 初始化每个导入的包，初始化顺序不是按照从上到下的导入顺序，而是按照解析的依赖关系，没有依赖的包最先初始化。

每个包首先初始化包作用域的常量和变量（常量优先于变量），然后执行包的 `init()` 函数。同一个包，甚至是同一个源文件可以有多个 `init()` 函数。`init()` 函数没有入参和返回值，不能被其他函数调用，同一个包内多个 `init()` 函数的执行顺序不作保证。

一句话总结： **`import –> const –> var –> init() –> main()`**

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

---

### 1 楼

根据包的导入顺序依次初始化包中的常量、变量、init()函数

### 5 楼

go的init函数非必要不使用

### 23 楼

init() 函数没有入参和返回值，不能被其他函数调用，同一个包内多个 init() 函数的执行顺序不作保证。

一句话总结： import –> const –> var –> init() –> main()

### 26 楼

init() 函数，是这个包被引用的时候执行，并且是末梢包的init先执行，同一个包可以有N 个init但是执行的顺序不保证，init不可以被我们编写的程序调用。

### 43 楼

奇怪为什么init函数能够被多次声明

---

### 资料

[五分钟理解 golang 的 init 函数](https://zhuanlan.zhihu.com/p/34211611)

</div>
</details>
