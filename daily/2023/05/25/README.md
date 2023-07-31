# Go 每日一题

> 今日（2023-05-25）的题目如下

以下代码能否编译？

```golang
package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {

	list = make(map[string]Student)

	student := Student{"Aceld"}

	list["student"] = student
	list["student"].Name = "LDB"

	fmt.Println(list["student"])
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

#### 结果

编译失败，`cannot assign to struct field list["student"].Name in map`

#### 分析

`map[string]Student` 的 value 是一个 Student 结构值，所以当`list["student"] = student`, 是一个值拷贝过程。而`list["student"]`则是一个值引用。那么值引用的特点是`只读`。所以对`list["student"].Name = "LDB"`的修改是不允许的。

#### 方法一：

```golang
package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {

	list = make(map[string]Student)

	student := Student{"Aceld"}

	list["student"] = student
	//list["student"].Name = "LDB"

    /*
        方法1:
    */
    tmpStudent := list["student"]
    tmpStudent.Name = "LDB"
    list["student"] = tmpStudent

	fmt.Println(list["student"])
}
```

其中

```golang
/**
方法1:
*/
tmpStudent := list["student"]
tmpStudent.Name = "LDB"
list["student"] = tmpStudent
```

是先做一次值拷贝，做出一个`tmpStudent副本`,然后修改该副本，然后再次发生一次值拷贝复制回去，`list["student"] = tmpStudent`,但是这种会在整体过程中发生 2 次结构体值拷贝，性能很差。

#### 方法二：

```golang
package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]*Student

func main() {

	list = make(map[string]*Student)

	student := Student{"Aceld"}

	list["student"] = &student
	list["student"].Name = "LDB"

	fmt.Println(list["student"])
}
```

我们将 map 的类型的 value 由 Student 值，改成 Student 指针。

```golang
var list map[string]*Student
```

这样，我们实际上每次修改的都是指针所指向的 Student 空间，指针本身是常指针，不能修改，`只读`属性，但是指向的 Student 是可以随便修改的，而且这里并不需要值拷贝。只是一个指针的赋值。

---

### 3 楼

看了答案，发现因为list["student"]是一个引用，所以是只读的。那如果把Student换成slice，却是可以修改的，这是为什么呢？ 如下代码 mapSlice["student"][1] = 4 没问题：

```golang
package main

import "fmt"

type Student struct {
    Name string
}

var mapStu map[string]Student
var mapSlice map[string][]int

func main() {
    mapStu = make(map[string]Student)
    stu := Student{"mark"}
    mapStu["student"] = stu
    //mapStu["student"].Name = "rudy"
    fmt.Println(mapStu)

    mapSlice = make(map[string][]int)
    s := []int{1, 2, 3}
    mapSlice["student"] = s
    mapSlice["student"][1] = 4
    fmt.Println(mapSlice)
}
```

### 4 楼

>   回复 3 楼

因为slice底层是一个*array


</div>
</details>
