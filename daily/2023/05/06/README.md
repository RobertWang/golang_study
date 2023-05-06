# Go 每日一题

> 今日（2023-05-06）的题目如下

以下代码是否能编译通过？

```golang
package main

import "fmt"

func main() {
	m := make(map[string]int)

	fmt.Println(&m["qcrao"])
}
```

<details>
<summary>答案解析：</summary>
<div>

这个问题，相当于问：可以对 map 的元素直接取地址吗？

以上代码编译报错：

>   ./main.go:8:14: cannot take the address of m["qcrao"]

即无法对 map 的 key 或 value 进行取址。

如果通过其他 hack 的方式，例如 unsafe.Pointer 等获取到了 key 或 value 的地址，也不能长期持有，因为一旦发生扩容，key 和 value 的位置就会改变，之前保存的地址也就失效了。

题目和解析来自：[https://golang.design/go-questions/map/element-address/](https://golang.design/go-questions/map/element-address/)

### 3楼

这种问题我称之为刻舟求剑。slice也有扩容问题，可是为什么不限制

```golang
func main() {
    type User struct {
        Age int
    }

    ss := []User{{Age: 23}}
    u1 := &ss[0]

    for i := 0; i < 10; i++ {
        ss = append(ss, User{Age: i})
    }

    u1.Age = 45 // 改改u1

    u2 := &ss[0]

    fmt.Println(u2)
    fmt.Println(ss[0])
}
```

### 6楼

>   回复 3 楼
slice 不用hash 扩容之后 虚拟地址相对位置没变地址，不知道可不可以这么理解


### 24楼

map 不能求地址 fmt.Println(&m["qcrao"])

### 25楼

>   考虑到map可以自动扩容，map中的数据元素的value位置可能在这一过程中发生变化，因此Go不允许获取map中value的地址，这个约束是在编译期间就生效的。


### 26楼

无法对 map 的 key 或 value 进行取址

### 27楼

这个问题，相当于问：可以对 map 的元素直接取地址吗？

即无法对 map 的 key 或 value 进行取址。

如果通过其他 hack 的方式，例如 unsafe.Pointer 等获取到了 key 或 value 的地址，也不能长期持有，因为一旦发生扩容，key 和 value 的位置就会改变，之前保存的地址也就失效了。


### 30楼

无法对 map 的 key 或 value 进行取址。

如果通过其他 hack 的方式，例如 unsafe.Pointer 等获取到了 key 或 value 的地址，也不能长期持有，因为一旦发生扩容，key 和 value 的位置就会改变，之前保存的地址也就失效了。

### 37楼

简单来说，就是直接对map的key或者value来进行取地址是不行的！就算通过其它的方式来获取到地址，也不能长期拥有，因为一旦进行扩容地址就会发生变化！


</div>
</details>
