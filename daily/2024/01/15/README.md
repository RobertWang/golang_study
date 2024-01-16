# 2024-01-15 练习题

> 来源 : 2024-01-15 [Python 每日一练-LeetCode-字符串练习-加密路径](https://www.bilibili.com/video/BV16C4y1k7K2/)

## 题目说明

假定一段路径记作字符串 path，其中以 "." 作为分隔符。现需将路径加密，加密方法为将 path 中的分隔符 "." 替换为空格 " "，请返回加密后的字符串。


## 示例说明

> 示例:
>
> 输入: path="a.aef.qerf.bb"
>
> 输出: "a aef gerf bb"


<details>
<summary style="cursor: pointer">🔑 参考：</summary>
<div>

## 分析

字符串替换即可

## 参考代码

### Golang 代码实现

```golang
package main

import (
	"fmt"
	"strings"
)

// 入口
func main() {
	data := "a.aef.qerf.bb"
	fmt.Println("初始数据", data)

	solution(data)
}

// 解决方案
func solution(password string) {
	encoded := strings.Replace(password, ".", " ", -1)
	fmt.Println("加密后的字符串为:", encoded)
}
```

### Python 代码实现

```python
path = input("enter path(.split):")
ls = path.split(".")
lls = " ".join(ls)
print(f"encoded path:{lls}")
```

用 split 方法将输入的字符串用 "." 进行分割
再使用 join 方法, 将分割生成的列表用空格进行合并
打印出合并后的加密路径即可

```
enter path(.split):a.aef.gerf.bb
encoded path:a aef gerf bb
```

</div>
</details>
