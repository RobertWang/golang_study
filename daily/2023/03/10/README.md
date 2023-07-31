# Go 每日一题

> 今日（2023-03-10）的题目如下

下面代码下划线处可以填入哪个选项？

```golang
func main() {
	var s1 []int
	var s2 = []int{}
	if __ == nil {
		fmt.Println("yes nil")
	}else{
		fmt.Println("no nil")
	}
}
```

- A. s1
- B. s2
- C. s1、s2 都可以

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：A。

知识点：nil 切片和空切片。nil 切片和 nil 相等，一般用来表示一个不存在的切片；空切片和 nil 不相等，表示一个空的集合。

---

### 3楼

为什么不是都可以，题目仅仅是可以填入，并没有说要走第一个分支

### 4楼

3 楼正解。都可以填， s2 == nil 返回的是 false, s1 == nil 返回的是 true


</div>
</details>
