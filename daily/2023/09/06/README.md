# Go 每日一题

> 今日（2023-09-06）的题目如下

下面这段代码能否通过编译，不能的话原因是什么；如果通过，输出什么。

```golang
func main() {
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

不能通过编译，new([]int) 之后的 list 是一个 `*[]int` 类型的指针，不能对指针执行 append 操作。可以使用 make() 初始化之后再用。同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。

## 1楼

正确写法

```golang
list := new([]int)
*list = append(*list, 1)
fmt.Println(*list)
// 或者
list := make([]int, 0)
list = append(list, 1)
fmt.Println(list)
```


</div>
</details>
