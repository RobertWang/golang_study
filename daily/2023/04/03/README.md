# Go 每日一题

> 今日（2023-04-03）的题目如下

下面这段代码输出什么？

```golang
func main() {
	m := map[int]string{0:"zero",1:"one"}
	for k,v := range m {
		fmt.Println(k,v)
	}
}
```

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：

```
0 zero
1 one
// 或者
1 one
0 zero
```

map 的输出是无序的。

---

### 28楼

map遍历无序，判断map中某元素存在与否

```golang
value, ok := map[key]
```


</div>
</details>
