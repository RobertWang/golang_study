# Go 每日一题

> 今日（2023-04-04）的题目如下

下面这段代码输出什么？

```golang
func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
```

<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：

```
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
```

程序执行到 main() 函数三行代码的时候，会先执行 calc() 函数的 b 参数，即：`calc("10",a,b)`，输出：10 1 2 3，得到值 3，因为 defer 定义的函数是延迟函数，故 `calc("1",1,3)` 会被延迟执行；

程序执行到第五行的时候，同样先执行 calc("20",a,b) 输出：20 0 2 2 得到值 2，同样将 calc("2",0,2) 延迟执行；

程序执行到末尾的时候，按照栈先进后出的方式依次执行：calc("2",0,2)，calc("1",1,3)，则就依次输出：2 0 2 2，1 1 3 4。

---

### 5 楼

```golang
package main

import (
    "fmt"
)

func main() {
    fmt.Println("main", inc())
}

func inc() int {
    t := &test{num: 0}
    defer t.Inc(3).Inc(2).Inc(1)
    fmt.Println("inc", t.num)
    return t.num
}

type test struct {
    num int
}

func (t *test) Inc(flag int) *test {
    t.num++
    fmt.Println("test", flag, t.num)
    return t
}
```

答案是

```
test 3 1
test 2 2
inc 2
test 1 3
main 2
```

我的这个测试代码更能够理解 defer 的逻辑。

在执行 `defer t.Inc(3).Inc(2).Inc(1)` 时，从左到右先计算了 `3`, `2` 这两个,最后的 `1` 才压栈。

根据结果可看到先执行 `return t.num`，因为 `main 2` 表示最终返回值是 2

所以 return 先执行，然后再执行之前压栈的 `defer t.Inc(1)`

</div>
</details>
