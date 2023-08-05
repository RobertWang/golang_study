# Go 每日一题

> 今日（2023-08-05）的题目如下

下面这段代码能否正常结束？

```golang
func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：不会出现死循环，能正常结束。

循环次数在循环开始前就已经确定，循环内改变切片的长度，不影响循环次数。

---

### 4 楼

最终v结果是[1 2 3 0 1 2] 循环次数是v的本来长度，感兴趣的可以去看看range的中间代码

### 6 楼

正常，range是副本循环

### 11 楼

即使切片是引用类型，但是循环次数是循环开始前就确定

</div>
</details>
