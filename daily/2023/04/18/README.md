# Go 每日一题

> 今日（2023-04-18）的题目如下

下面这段代码输出什么？

```golang
func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(f(3))
}
```

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：7。

---

### 8 楼

应该是 defer 注册延迟执行函数时, f 未初始化, 为 nil, f()会 panic, r += 2 不会执行, 若把 defer f()与 f = func() {r += 2}换一下位置, 值就是 9 了

### 9 楼

> 6 楼

```golang
// 先执行这一段代码, 然后 r 被赋值为 4
return n+1

// 然后执行，但是会触发panic
defer f()

// 最后执行，r 再加3 r=7,然后recover(),最后函数返回 7
defer func() {
	r += n
	recover()
}()
```

### 12 楼

个人见解：

1. 考察了 defer 的调用，通过上面的`defer f()`可知，f()是没有定义的？(存疑)所以报了 panic，而在`defer func()`里面的`recover()`处理了这个 panic，保证了整个程序能够执行完
2. 然后看 f()函数返回的值为 r，而不是 n+1，这里考察的是函数的返回这一块的知识，所以最后的结果不是`return n+1`而是`n+1+n`的结果（第三点详细解释）
3. 为什么会等于 7？看看最`f()`函数的最后一行为`return n+1`,也就是`return 3+1`,函数`f()`的返回参数为`r int`，所以`r = 3 + 1`,再到`defer func()`的`r+=n`也就是`r = n + 1 + n`即`r = 3 + 1 + 3`，然后返回 r 的值，就是 7

### 15 楼

可以参考 [https://studygolang.com/articles/27408](https://studygolang.com/articles/27408) 对于 defer 的知识点

### 40 楼

解释如下：

f(3)调用时，首先执行 return n + 1，即返回值为 4。 接着执行 defer f()，即执行 f 函数，r 的值变成了 6。 执行 defer func() {...}()，其中的 recover()函数可以使程序从 panic 状态中恢复过来，但是并没有发生 panic，所以这个函数不做任何操作。 因此最终的返回值是 4 + 2 + 1 = 7。

</div>
</details>
