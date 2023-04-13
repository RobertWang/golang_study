# Go 每日一题

> 今日（2023-04-11）的题目如下

下面代码输出什么？

```golang
type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

func main() {
	m["foo"].x = 4
	fmt.Println(m["foo"].x)
}
```

- A. 4
- B. compilation error

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：B

类似于 X=Y 的赋值操作,必须知道 X 的地址,才能够将 Y 的值赋给 X,但 go 中的 map 的 value 本身是不可寻地址。

</div>
</details>
