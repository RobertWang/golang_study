# Go 每日一题

> 今日（2023-08-07）的题目如下

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
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：7。

---

### 9 楼

应该是 defer 注册延迟执行函数时, f 未初始化, 为 nil, f()会 panic, r += 2 不会执行, 若把 defer f()与 f = func() {r += 2}换一下位置, 值就是 9 了

### 12 楼

个人见解：

- 考察了 defer 的调用，通过上面的 defer f()可知，f()是没有定义的？(存疑)所以报了 panic，而在 defer func()里面的 recover()处理了这个 panic，保证了整个程序能够执行完
- 然后看 f()函数返回的值为 r，而不是 n+1，这里考察的是函数的返回这一块的知识，所以最后的结果不是 return n+1 而是 n+1+n 的结果（第三点详细解释）
- 为什么会等于 7？看看最 f()函数的最后一行为 return n+1,也就是 return 3+1,函数 f()的返回参数为 r int，所以 r = 3 + 1,再到 defer func()的 r+=n 也就是 r = n + 1 + n 即 r = 3 + 1 + 3，然后返回 r 的值，就是 7

### 15 楼

可以参考 [https://studygolang.com/articles/27408](https://studygolang.com/articles/27408) 对于 defer 的知识点

### 40 楼

这里在我的博客里进行了讲解，不懂的 uu 可以参考一下，若有不对的地方，欢迎指出 解释如下：

首先使用 defer 关键字注册了一个匿名函数，然后这个匿名函数在函数 f 返回时执行。在这个匿名函数里，使用了 recover()，这意味着它可以恢复 panic。 接着定义了一个变量 f，类型为 func()，这里由于只声明了，但是没有定义，故变量 f 是一个 nil 函数。 然后使用 defer 关键字将 f 变量注册成延迟函数，这个延迟函数在函数 f 返回时会执行，但这个匿名函数是一个 nil 函数，因此在执行这个延迟函数时会触发 panic 接下来是对变量 f 的定义 return n + 1 此时，返回值变量 r = n + 1，接着执行 defer 注册的延迟函数，因为 defer 函数的执行顺序是先进后出的，故先执行变量 f，但由于这里注册的是一个 nil 函数，因此触发 panic，接着执行最开始注册的匿名函数，此时 r = n + 1 + n，遇到了 recover()，所以恢复了 panic，将 r 的值返回 最后返回给主函数的值 r = n + 1 + n = 7

</div>
</details>
