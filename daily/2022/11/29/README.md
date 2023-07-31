# Go每日一题

> 今日（2022-11-29）的题目如下

下面代码输出什么？

```golang
package main

import (
	"fmt"
)

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
	var a A = Work{3}
	// 类型断言
	// http://c.biancheng.net/view/4281.html
	s := a.(Work)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}
```

- A. 13 23
- B. compilation error


<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

```
13
23
```

参考答案及解析：A。

### 知识点：类型断言。

这道题可以和第 15 天的第三题 和第 16 天的第三题结合起来看。

</div>
</details>
