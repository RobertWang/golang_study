# Go 每日一题

> 今日（2023-03-19）的题目如下

下面代码输出什么？

```golang
func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}
```

- A. 1 1
- B. 0 1
- C. 1 0
- D. 0 0


<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：B。

知识点：defer、返回值。注意一下，increaseA() 的返回参数是匿名，increaseB() 是具名。关于 defer 与返回值的知识点，后面我会写篇文章详细分析，到时候可以看下文章的讲解。

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
