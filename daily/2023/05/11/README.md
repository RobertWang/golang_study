# Go 每日一题

> 今日（2023-05-11）的题目如下

下面这段代码能否通过编译，不能的话原因是什么；如果通过，输出什么。

```golang
func main() {
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}
```

<details>
<summary>答案解析：</summary>
<div>

不能通过编译，new([]int) 之后的 list 是一个 `*[]int` 类型的指针，不能对指针执行 append 操作。可以使用 make() 初始化之后再用。同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。

---

### 1 楼

new 其实在 go 里面没必要提供的，应该是为了反射，才提供 new 关键字

### 2 楼

new 在数据结构那方面应用还是很多的, 比如构建树那些东西

### 8 楼

> 回复 1 楼

半吊子才敢大放厥词。new 到这儿都成关键字了。 自己想不到就说人没用？什么混蛋逻辑？

>   As an exception to the addressability requirement, x may also be a (possibly parenthesized) composite literal.

如果不是这儿开了个口子，还想addressable？哪儿凉快哪儿哭去吧

### 18 楼

new 返回的为指针，这里的话，返回切片的指针，不能直接用于 append 操作，应该解引用 list := *new([]int)


</div>
</details>
