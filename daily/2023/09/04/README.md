# Go 每日一题

> 今日（2023-09-04）的题目如下

下面代码是否可以编译通过？为什么？

```golang
package main

import "fmt"

func main() {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

编译不通过。

>     ./prog.go:31:9: invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)

考点：结构体比较

>     **结构体比较规则注意1**：只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。

比如：

```golang
sn1 := struct {
	age  int
	name string
}{age: 11, name: "qq"}

sn3:= struct {
    name string
    age  int
}{age:11, name:"qq"}
```

`sn3`与`sn1`就不是相同的结构体了，不能比较。

>     **结构体比较规则注意2**：结构体是相同的，但是结构体属性中有不可以比较的类型，如`map`,`slice`，则结构体不能用`==`比较。

可以使用 reflect.DeepEqual 进行比较

```golang
if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
} else {
		fmt.Println("sm1 != sm2")
}
```

---

### 3 楼

map 实例是不可以直接比较的，如果非要比较的话，可以使用 reflect.DeepEqual(). go 中表明了可比较的类型：

- 布尔值
- 数字
- 字符串
- 指针
- 通道
- 接口类型
- 结构体

只包含上述类型的数组。

### 16 楼

- 布尔 是 不能与 nil 比较
- 整型 是 不能与 nil 比较
- 浮点数 是 不能与 nil 比较
- 复数 是 不能与 nil 比较
- 字符串 是 不能与 nil 比较
- 指针值 是 两个指针指向同一个变量，则这两个指针相等，或者两个指针同为 nil，它们也相等。指针值可以与 nil 比较。
- 通道值 是 两个通道是由同一个 make 创建的，或者两个通道值都为 nil，那么它们是相等，不同类型的通道都是 nil 也不可比较
- 接口值 是 如果两个接口值的动态值和动态类型都相等，或者两个接口值都为 nil，那么它们是相等的。接口值可以与 nil 进行比较。
- struct 是/否 如果 struct 中所有的字段都是可比较的，那么两个 struct 是可比较的。如果 struct 对应的非空白字段相等，则它们相等。struct 不能与 nil 比较。
- 数组 是/否 如果数组中的元素类型是可比的，则数组也是可比较的。如果数组中对应的元素都相等元素顺序必须一样，那么两个数组是相等的。数组不能与 nil 比较
- map 否 只能与 nil 比较
- slice 否 只能与 nil 比较
- func 否 只能与 nil 比较

### 24 楼

1. 只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关
2. 结构体属性中有不可以比较的类型，如 map,slice，则结构体不能用==比较, reflect.DeepEqual

</div>
</details>
