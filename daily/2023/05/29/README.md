# Go 每日一题

> 今日（2023-05-29）的题目如下

以下代码能通过编译吗？为什么？

```golang
package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "love" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = Student{}
	think := "love"
	fmt.Println(peo.Speak(think))
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

考点 : 继承与多态的特点

在 golang 中对多态的特点体现从语法上并不是很明显。

我们知道发生多态的几个要素：

1. 有 interface 接口，并且有接口定义的方法。
2. 有子类去重写 interface 的接口。
3. 有父类指针指向子类的具体对象

那么，满足上述 3 个条件，就可以产生多态效果，就是，父类指针可以调用子类的具体方法。

所以上述代码报错的地方在 `var peo People = Student{}` 这条语句， `Student{}` 已经重写了父类 `People{}` 中的 `Speak(string) string` 方法，那么只需要用父类指针指向子类对象即可。（Go 中不叫父类，这里是为了好理解）

所以应该改成 `var peo People = &Student{}` 即可编译通过。（People 为 interface 类型，就是指针类型）

---

### 12 楼

我觉得强制用java等基于类的面向对象的思维解释golang是很奇怪且不正确的.

对于这个问题, 我觉得编译器的报错已经很清晰地回答了:

>	Student does not implement People (Speak method has pointer receiver)

也就是说, `Student` 不是 `People` 类型, 但是 `*Student` 是. 就这么简单. 只是当写下

```golang
peo := Student{}
peo.Speak("hello")
```

时, 编译器会隐式替换为 `(&peo).Speak("Hello")` 而已, 给人以 `Student` 也有 `func Speak(string) string` 方法的错觉.

ref: [https://books.studygolang.com/gopl-zh/ch6/ch6-02.html](https://books.studygolang.com/gopl-zh/ch6/ch6-02.html)

p.s.

```golang

// ...


func main() {
    var p2 interface{} = Student{}
    var p3 interface{} = &Student{}
    if _, ok := p2.(People); ok {
    fmt.Println("p2 is of type People")
    } else if _, ok := p3.(People); ok {
    fmt.Println("p3 is of type People")
    }
}
```

output:

```
p3 is of type People
```

### 14 楼

那么如果创建Student{}或者&Student{}会自动取引用和解引用,我估计这里是接口类型不能自动取引用和解引用到子类特定的对象去引用方法,父类对象指向子类的对应方法的类型,必须要是子类接收者完全对应的类型

### 15 楼

首先, golang没有父类和子类的概念, 任何interface类型不是任何结构的所谓"父类". 其次, interface只认哪个类型是接收器, 类型和类型指针是两个类型, 分到两个接收器. 我觉得 [https://books.studygolang.com/gopl-zh/ch6/ch6.html](https://books.studygolang.com/gopl-zh/ch6/ch6.html) 和 [https://books.studygolang.com/gopl-zh/ch7/ch7.html](https://books.studygolang.com/gopl-zh/ch7/ch7.html) 这两章已经把很多东西都讲清楚了.

### 18 楼

不能通过编译。 var peo People = Student{} peo是一个对象，实现的方法是针对指针实现的。 var peo People = &Student{} 就OK了。

### 40 楼

还有一个解法：将重写的 `Speak` 方法的接收对象改为非指针对象 `func (stu Student) Speak(think string) (talk string) {`

### 41 楼

方法接受者是t *T，赋值只能是指针； 方法接受者是t T，赋值可以是t / &t



</div>
</details>
