# Go 每日一题

> 今日（2023-03-20）的题目如下

下面代码输出什么？

```golang
type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	var a A = Work{3}
	s := a.(Work)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}
```

- A. 23 13
- B. compilation error


<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：A。

知识点：类型断言。这道题可以和第 15 天的第三题 和第 16 天的第三题结合起来看。

---

### 7楼

错，对，对，错 3是个知识点：当多值赋值时，:= 左边的变量无论声明与否都可以

### 10楼

2、3 对；赋值多个变量，只要有一个变量时新的，就可以用“:=”

### 11楼

“当多值赋值时，:= 左边的变量无论声明与否都可以” 小编在吗？应该是，“至少有一个变量是新的”

### 16楼

```golang
var a, b int
a, b := 1, 2
```

多赋值，至少有一个变量是新的


</div>
</details>
