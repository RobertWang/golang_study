# Go 每日一题

> 今日（2023-06-13）的题目如下

以下代码能否通过编译？

```golang
package main

import (
	"fmt"
)

type worker interface {
	work()
}

type person struct {
	name string
	worker
}

func main() {
	var w worker = person{}
	fmt.Println(w)
}
```

- A: 能
- B: 不能
- C: 不知道


<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

这里答错的原因在于 worker 是一个接口，如果是一个普通的类型，相信大家会答对。一个结构体竟然可以嵌入一个接口？！

### 01

我们都知晓 Go 没有继承，但可以通过内嵌类型模拟部分继承的功能。大家要记住，接口也是类型，自然也将它作为嵌入类型。如果题目的 person 接口体改为：

```golang
type person struct {
	name string
	worker worker
}
```

相信会有更多人答对，这和嵌入类型唯一的区别在于是否显示指定了字段名，其他并无区别。

将接口作为嵌入类型可能让人感觉有些奇怪：那这个类型不是默认就实现了这个接口？！确实是这样的，所以才有了题目中这一句能编译通过：

```golang
var w worker = person{}
```

只不过，因为实例化 person 时，没有给 worker 指定值，因此 person 中的 worker 是 nil，调用它的话会报错，但编译是没问题的。

### 02

有人可能要问，嵌入接口有实际用途吗？我找一个标准库中的例子。

在 sort 包中，有一个接口：Interface

```golang
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```

这是用于排序的。还有另外一个结构体：[reverse](https://docs.studygolang.com/src/sort/sort.go?s=7078:7105#L239)

```golang
type reverse struct {
	// This embedded Interface permits Reverse to use the methods of
	// another Interface implementation.
	Interface
}
```

它就内嵌了一个 Interface，用于排序的反转。

而内嵌接口的关键在于如何给这个内嵌的接口赋值。sort 包有一个函数：Reverse

```golang
func Reverse(data Interface) Interface {
	return &reverse{data}
}
```

其中实例化 reverse 时，直接通过传递的 Interface 实例赋值给 reverse 的内嵌接口，然后 reverse 类型可以有选择的重新实现内嵌的 Interface 的方法。比如 Less 方法：

```golang
func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}
```

### 03

回到上面的题目，如果我们通过实例化的 w 调用 work 方法会报错：

```golang
var w worker = person{}
w.work()	// panic
```

和上面 reverse 类似，你需要给 person 中的 worker 实例化，也就是需要一个实现了 worker 接口的类型实例。比如：

```golang
type student struct{
    name string
}

func (s student) work() {
    fmt.Println("I am ", s.name, ", I am learning")
}
```

然后这样实例化 person：

```golang
var w worker = person{worker: student{"polarisxu"}}
```

你掌握了吗？

答案解析来自：[https://polarisxu.studygolang.com/posts/go/action/weekly-question-analysis-embed-interface/](https://polarisxu.studygolang.com/posts/go/action/weekly-question-analysis-embed-interface/)

---

### 3 楼

struct 中竟然能直接嵌入interface，学习到了。同时了解到了一种struct中嵌入interface的使用场景。

### 5 楼

struct中内嵌接口 和 实例化接口。关键在于 内嵌类型 实现了接口。一个类似继承，一个就是普通的组合。worker woker写法会让结构体增加一级

### 26 楼

Embedding Interfaces in Go (Golang)
- [https://golangbyexample.com/embedding-interfaces-go/](https://golangbyexample.com/embedding-interfaces-go/)
- [https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/](https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/)


</div>
</details>
