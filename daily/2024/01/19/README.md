# 2024-01-19 练习题

> 来源 : 2024-01-19 [Python 每日一练-LeetCode-字符串练习-罗马数字转换为整数](https://www.bilibili.com/video/BV125411e7Ag/)

## 题目说明

将一个罗马数字转换为整数。

## 示例说明

> 例如：
>
> - 输入：s="III"
> - 输出：3
>
> - 输入：s="IX"
> - 输出：9
>
> - 输入：s="IV"
> - 输出：4
>
> - 输入：s="LVIII"
> - 输出：58
> - 解释：L=50，V=5，III=3。
>
> - 输入：s="MCMXCIV"
> - 输出：1994
> - 解释：M=1000，CM=900，XC=90，IV=4

### 罗马数字对照表

| 字符 | 数值 |
| ---- | ---- |
| I    | 1    |
| V    | 5    |
| X    | 10   |
| L    | 50   |
| C    | 100  |
| D    | 500  |
| M    | 1000 |

### 罗马数字表示规则说明

罗马数字 2 写做 II，即为两个并列的 1。
12 写做 XII，即为 X+II。27 写做 XXVII，即为 XX+V+II。

通常情况下，罗马数字中小的数字在大的数字的右边。

但也存在特例：

> 例如 4 不写做 IIII，而是 IV。
>
> 数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4
> 同样地，数字 9 表示为 IX。

这个特殊的规则只适用于以下六种情况：

- ① 1 可以放在 V(5)、X(10)的左边，来表示 4 和 9。
- ② X 可以放在 L(50)、C(100)的左边，来表示 40 和 90。
- ③ C 可以放在 D(500)、M(1000)的左边来表示 400 和 900。

<details>
<summary style="cursor: pointer">🔑 参考：</summary>
<div>

## 分析

按照转换规律，正常情况下遍历每个字符，罗马字符从小到大排列，累加上字符对应的数值即可。

但对于六种特殊情况，当数值小于后面一个字符的数值时，前一个小的数值要被减去。

所以通过 enumeratei 遍历出罗马数字的字符和索引。

当字符的索引不是最后一个字符，且对应的数值小于其后一个索引对应数值时减去该数值。用 resulti 存放转换后的数值

## 参考代码

### Golang 代码实现

```golang
package main

import (
	"fmt"
)


// 实例 1: "LVIIIMCMXCIV" // 2050
// 入口
func main() {
	data := "LVIIIMCMXCIV"
	fmt.Println("初始数据", data)

	solution(data)
}

// 解决方案
func solution(roma_num string) {
	result := 0
	length := len(roma_num)
	symbol := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}
	for i, char := range roma_num {
		if i < length-1 && symbol[string(char)] < symbol[string(roma_num[i+1])] {
			result -= symbol[string(char)]
		} else {
			result += symbol[string(char)]
		}	
	}
	fmt.Println(roma_num, "转换为整数为", result)
}
```

### Python 代码实现

```python
s = "LVIIIMCMXCIV"
# 转换为整数为: 2050

symbol = {
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000
}
result = 0
n = len(s)
for i, ch in enumerate(s):
    value = symbol[ch]
    # 如果这个罗马数字既不是最后一个：又小于后面二个数守
    # 即属于六种特殊情况之裴减去这个值
    if i < n-1 and value < symbol[s[i+1]]:
        result -= value
    else:
        result += value
print(f"{s}转换为整数为：{result}")
```

</div>
</details>
