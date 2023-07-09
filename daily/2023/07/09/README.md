# Go 每日一题

> 今日（2023-07-09）的题目如下

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

### 9 楼

B. 0 1 ；
`func increaseA() int {` ，返回值 i=0 的时候已经绑定到返回值里里，defer 改 i 没用了。 `func increaseB() (r int) {` 先把返回变量 r 设为 0，defer 又把 r 改为 1.

### 15 楼

return 语句是把 return 后面的值赋给返回值，但由于 increaseB 中有 return 变量 r，所以在 defer 里还可以对变量 r 进行再更新，所以返回的 r 为 1

</div>
</details>
