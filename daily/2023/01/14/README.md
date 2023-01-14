# Go每日一题

> 今日（2023-01-14） 的题目如下

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

```
./main.go:8:14: cannot take the address of m["qcrao"]
```

即无法对 map 的 key 或 value 进行取址。

如果通过其他 hack 的方式，例如 unsafe.Pointer 等获取到了 key 或 value 的地址，也不能长期持有，因为一旦发生扩容，key 和 value 的位置就会改变，之前保存的地址也就失效了。

题目和解析来自：https://golang.design/go-questions/map/element-address/

</div>
</details>