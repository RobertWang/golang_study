# Go每日一题

> 今日（2023-01-29）的题目如下

下面这段代码能否通过编译，如果可以，输出什么？

```golang
var(
	size := 1024
	max_size = size*2
)

func main() {
	fmt.Println(size,max_size)
}
```


<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案：不能通过编译。

参考解析：这道题的主要知识点是变量声明的简短模式，形如：x := 100.
但这种声明方式有限制：

- 必须使用显示初始化；
- 不能提供数据类型，编译器会自动推导；
- 只能在函数内部使用简短模式；


### 22楼

:=赋值：函数内部，不能指定类型，需要系统推导，显示初始化


</div>
</details>