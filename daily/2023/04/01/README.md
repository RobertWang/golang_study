# Go 每日一题

> 今日（2023-04-01）的题目如下

下面这段代码输出什么？为什么？

```golang
func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1)
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：

```
[1 2 4]
[1 2 4]
```

我们已经知道，golang 中切片底层的数据结构是数组。当使用 s1[1:] 获得切片 s2，和 s1 共享同一个底层数组，这会导致 s2[1] = 4 语句影响 s1。

而 append 操作会导致底层数组扩容，生成新的数组，因此追加数据后的 s2 不会影响 s1。

但是为什么对 s2 赋值后影响的却是 s1 的第三个元素呢？这是因为切片 s2 是从数组的第二个元素开始，s2 索引为 1 的元素对应的是 s1 索引为 2 的元素。

---

### 7楼

这个更难，有人回答一下吗？ 

```golang
func main() {
    s1 := []int{1, 2, 3, 4, 9, 9, 9, 9}
    s2 := s1[1:4]
    s2[1] = 4
    fmt.Println(s1)
    s2 = append(s2, 5, 6, 7)
    fmt.Println(s1)
}
```

### 25楼

```golang
func Test0301(t *testing.T) {
    s1 := []int{1, 2, 3, 4, 9, 9, 9, 9}
    s2 := s1[1:4]            //[2,3,4]
    s2[1] = 4                //[2,4,4]
    fmt.Println(s1)          //[1,2,4,4,9,9,9,9]
    s2 = append(s2, 5, 6, 7) //[2,4,4,5,6,7]
    fmt.Println(s1)          //[1,2,4,4,5,6,7,9]
}
```

</div>
</details>
