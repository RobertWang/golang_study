# Go 每日一题

> 今日（2023-07-02）的题目如下

下面这段代码输出什么？

```golang
type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}
```

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：13 23。

知识点：接口。一种类型实现多个接口，结构体 Work 分别实现了接口 A、B，所以接口变量 a、b 调用各自的方法 ShowA() 和 ShowB()，输出 13、23。


---

### 17 楼

一种类型实现多个接口，结构体 Work 分别实现了接口 A、B，所以接口变量 a、b 调用各自的方法 ShowA() 和 ShowB()，输出 13、23

### 18 楼

13，23，就算实现方法参数是w *Work答案依然不变，因为实现方法本身只是返回运算结果，并没有修改传入参数的值


</div>
</details>
