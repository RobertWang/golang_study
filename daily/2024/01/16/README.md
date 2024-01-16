# 2024-01-16 练习题

> 来源 : 2024-01-16 [Python 每日一练-LeetCode-字符串练习-FizzBuzz 表达](https://www.bilibili.com/video/BV1qi4y1i7Dw/)

## 题目说明

给你一个整数,找出以 1 到 n 各个整数的 FizzBuzz 表示, 并用字符串数组 answer (下标从 1 开始) 返回结果, 其中:

- answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数
- answer[i] == "Fizz" 如果 i 是 3 的倍数.
- answer[i] == "Buzz" 如果 i 是 5 的倍数.
- answer[i] == ⅰ（以字符串形式） 如果上述条件全不满足.

## 示例说明

> 输入: n=3
>
> 输出: ["1","2","Fizz"]
>
> 输入: n=5
>
> 输出: ["1","2","Fizz","4","Buzz"]
>
> 输入: n=15
>
> 输出: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11""Fizz","13","14""FizzBuzz"]

<details>
<summary style="cursor: pointer">🔑 参考：</summary>
<div>

## 分析

## 参考代码

### Golang 代码实现

```golang
package main

import (
	"fmt"
	"strings"
)

// 实例 1: 5
// 入口
func main() {
	data := 5
	fmt.Println("初始数据", data)

	solution(data)
}

// 解决方案
func solution(num int) {
	var series []interface{}
	var i int
	var item interface{}

	for i = 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			item = "FizzBuzz"
		} else if i%3 == 0 {
			item = "Fizz"
		} else if i%5 == 0 {
			item = "Buzz"
		} else {
			item = i
		}
		series = append(series, item)
	}
	fmt.Println("结果序列为:", series)
}
```

### Python 代码实现

```python
n = int(input("enter a int number:")
ls = []
for i in range(1, n+1):
    if i % 3 == 0 and i % 5 == 0:
        ls.append("FizzBuzz")
    elif i % 3 == 0:
        ls.append("Fizz")
    elif i % 5 == 0:
        ls.append("Buzz")
    else:
        ls.append(str(i))
print(ls)
```

题目要求编号从 1~n
能被 3 和 5 整除的编号处写入 “FizzBuzz”
能被 3 整除的编号处写入 Fizz
能被 5 整除的编号处写入 Buzz
其余写入编号的字符串

```
enter a int number: 5
["1","2","Fizz","4","Buzz"]
```

</div>
</details>
