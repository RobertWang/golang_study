# Go 每日一题

> 今日（2023-05-27）的题目如下

下面这段代码能否编译通过？如果可以，输出什么？

```golang
const (
	x = iota
	_
	y
	z = "zz"
	k 
	p = iota
)

func main()  {
	fmt.Println(x,y,z,k,p)
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案：编译通过，输出：0 2 zz zz 5

参考解析：知识点 iota。

参考 [https://www.cnblogs.com/zsy/p/5370052.html](https://www.cnblogs.com/zsy/p/5370052.html)

---

### 1 楼

iota为Go中的常量计数器，同时会在const关键字出现时被重置为0。可以把iota理解为const语句块中的行索引，所以p的值为5


</div>
</details>
