# Go 每日一题

> 今日（2023-09-08）的题目如下

写出程序运行的结果：

```golang
package main

import (
    "fmt"
)

func main(){
    s := make([]int, 10)
    s = append(s, 1, 2, 3)
    fmt.Println(s)
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

结果

```golang
[0 0 0 0 0 0 0 0 0 0 1 2 3]
```

考点

切片追加, make 初始化均为 0

</div>
</details>
