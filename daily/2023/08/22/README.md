# Go 每日一题

> 今日（2023-08-22）的题目如下

下面两段代码输出什么。

```golang
// 1.
func main() {
    s := make([]int, 5)
    s = append(s, 1, 2, 3)
    fmt.Println(s)
}

// 2.
func main() {
	s := make([]int,0)
	s = append(s,1,2,3,4)
	fmt.Println(s)
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

```
代码 1 输出：[0 0 0 0 0 1 2 3]
代码 2 输出：[1 2 3 4]
```

参考解析：这道题考的是使用 append 向 slice 添加元素，第一段代码常见的错误是 [1 2 3]，需要注意。

---

### 27 楼

```golang
s := make([]int, 5)
```

这里是初始化切片 s 的长度为 5（当然容量也为 5），其中 int 类型的零值当然是 0 啦，所以 s 是[0 0 0 0 0]

### 31 楼

make(T[], len, cap) 创建切片，其中：

- T：切片的元素类型
- len：切片的实际长度
- cap：切片的最大容量

注意：当 cap 不传值的话，默认 len=cap。

举例：s := make(int[], 5)。此时 s 的值为 [0 0 0 0 0]

</div>
</details>
