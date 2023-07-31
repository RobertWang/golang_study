# Go 每日一题

> 今日（2023-03-21）的题目如下

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
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：1 5 1。

知识点：defer、返回值。

---

### 9楼

- f1() =1，return 把r设成0，然后defer把r改为1 ；
- f2() =5，return 把r设成5，然后defer改的是t，不影响返回值 ；
- f3() =5，return 把r设成1，然后defer把r改为r+5，但是用的r是defer设定时的r，=0；（靠，是1，r+5的r不是外面的r）


### 27楼

f(3)defer内部的r非返回值r



</div>
</details>
