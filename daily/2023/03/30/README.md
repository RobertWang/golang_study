# Go 每日一题

> 今日（2023-03-31）的题目如下

return 之后的 defer 语句会执行吗，下面这段代码输出什么？

```golang
var a bool = true
func main() {
	defer func(){
		fmt.Println("1")
	}()
	if a == true {
		fmt.Println("2")
		return
	}
	defer func(){
		fmt.Println("3")
	}()
}
```


<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：

```
2
1
```

defer 关键字后面的函数或者方法想要执行必须先注册，return 之后的 defer 是不能注册的， 也就不能执行后面的函数或方法。

Reference: [https://studygolang.com/topics/9967](https://studygolang.com/topics/9967)

</div>
</details>
 