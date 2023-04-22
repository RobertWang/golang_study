# Go 每日一题

> 今日（2023-04-23）的题目如下

下面代码里的 counter 的输出值？

```golang
func main() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}
```

- A. 2
- B. 3
- C. 2 或 3

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：C。

for range map 是无序的，如果第一次循环到 A，则输出 3；否则输出 2。

---

### 4 楼

map 是在哈希散列表，当使用 for range 时是无序的

### 12 楼

range map 未指定遍历次数 [https://cloud.tencent.com/developer/article/1876243](https://cloud.tencent.com/developer/article/1876243)

### 15 楼

for range 遍历切片时，循环次数是确定的，因此在循环体内改变切片结构对循环结果不影响；而 for range 遍历 map 时，循环次数是不确定的，循环体内对 map 的修改会在结果上反映出来

</div>
</details>
