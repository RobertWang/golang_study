# Go 每日一题

> 今日（2023-08-20）的题目如下

下面这段代码输出的内容

```golang
package main

import (
    "fmt"
)

func main() {
    defer_call()
}

func defer_call() {
    defer func() { fmt.Println("打印前") }()
    defer func() { fmt.Println("打印中") }()
    defer func() { fmt.Println("打印后") }()

    panic("触发异常")
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

```
打印后
打印中
打印前
panic: 触发异常
```

解析：defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行 panic。

---

### 36 楼

哈哈哈，可以这么变化下，感觉会误导很多人：

```golang
package main

import "fmt"

func defer_call() {
    defer func() {
        if p := recover(); p != nil {
            fmt.Println(p)
        }
    }()
    defer func() {
        fmt.Println("打印前")
    }()
    defer func() {
        fmt.Println("打印中")
    }()
    defer func() {
        fmt.Println("打印后")
    }()

    panic("触发异常")
}

func main() {
    defer_call()
}
```

</div>
</details>
