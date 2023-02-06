# Go每日一题

> 今日（2023-02-06）的题目如下

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
<summary>答案解析：</summary>
<div>

继承与多态的特点
在 golang 中对多态的特点体现从语法上并不是很明显。

我们知道发生多态的几个要素：

- 1、有interface接口，并且有接口定义的方法。
- 2、有子类去重写interface的接口。
- 3、有父类指针指向子类的具体对象

那么，满足上述 3 个条件，就可以产生多态效果，就是，父类指针可以调用子类的具体方法。

所以上述代码报错的地方在 `var peo People = Student{}` 这条语句， `Student{}` 已经重写了父类 `People{}` 中的 `Speak(string) string` 方法，那么只需要用父类指针指向子类对象即可。（Go 中不叫父类，这里是为了好理解）

所以应该改成 `var peo People = &Student{}` 即可编译通过。（People 为 interface 类型，就是指针类型）

### 15楼

首先, golang没有父类和子类的概念, 任何interface类型不是任何结构的所谓"父类". 其次, interface只认哪个类型是接收器, 类型和类型指针是两个类型, 分到两个接收器. 我觉得[https://books.studygolang.com/gopl-zh/ch6/ch6.html](https://books.studygolang.com/gopl-zh/ch6/ch6.html) 和 [https://books.studygolang.com/gopl-zh/ch7/ch7.html](https://books.studygolang.com/gopl-zh/ch7/ch7.html) 这两章已经把很多东西都讲清楚了.

### 24楼

编译不通过，实现Speak方法的属主是指针，所以要var peo People = &Student{}

### 25楼

func (stu Student) Speak 去掉引用类型就可以编译通过了。



</div>
</details>