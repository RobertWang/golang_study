# Go 每日一题

> 今日（2023-07-11）的题目如下

f1()、f2()、f3() 函数分别返回什么？

```golang
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}


func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}


func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
```


<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：1 5 1。

知识点：defer、返回值。


</div>
</details>
