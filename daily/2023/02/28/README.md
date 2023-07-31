# Go 每日一题

> 今日（2023-02-28）的题目如下

下面这段代码输出什么？

```golang
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowB()
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：teacher showB。

知识点：结构体嵌套。在嵌套结构体中，People 称为内部类型，Teacher 称为外部类型；通过嵌套，内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。此外，外部类型还可以定义自己的属性和方法，甚至可以定义与内部相同的方法，这样内部类型的方法就会被“屏蔽”。这个例子中的 ShowB() 就是同名方法。

---

### 1 楼

子类和父类方法名相同，子类的优先

### 8 楼

不要在 golang 里面谈父类子类和继承。类比都不要类比。纯属误导人！！！

golang 没有类与继承。有接口，内嵌，组合。

与面向对象的类对应的概念。不是 struct，是 interface！！

那些觉得 struct 对应类的。我问你，是不是觉得只有 struct 才能有方法？

事实上 golang 里面任何类型都可以有方法。接收器的类型远不止是 struct。甚至连 int 都可以是接收器。

### 9 楼

结构体嵌套。在嵌套结构体中，People 称为内部类型，Teacher 称为外部类型；通过嵌套，内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。此外，外部类型还可以定义自己的属性和方法，甚至可以定义与内部相同的方法，这样内部类型的方法就会被“屏蔽”。这个例子中的 ShowB() 就是同名方法。

### 11 楼

你只学 GO 语言吗？面向对象的概念而已，很多语言都通用。

### 15 楼

你可以看看这个视频，里面有讲到如何查看方法关联的情况

https://www.bilibili.com/video/BV1ZP4y1H7Gk?spm_id_from=333.999.0.0

### 26 楼

结构体的匿名嵌套本质就是继承，如果存在相同的方法则调用字段的方法，必须显性调用。不相同可省略字段直接调方法

```golang
type person struct{}  //mehtod showa  showb
type teacher struct{ person}  //method showb showc

t := techer{}
t.showb
t.showc

t.person.showb
t.showa   // t.person.showa
```

### 34 楼

【卷卷面试题 02：组合式继承】 https://www.bilibili.com/video/BV1ZP4y1H7Gk

</div>
</details>
