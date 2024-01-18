# 2024-01-17 练习题

> 来源 : 2024-01-17 [Python 每日一练-LeetCode-字符串练习-计算日期天数](https://www.bilibili.com/video/BV1294y1N7id/)

## 题目说明

给你一个字符串 date，按 YYYY-MM-DD 格式表示一个现行公元纪年法日期。

返回该日期是当年的第几天。


## 示例说明

> 示例1:
>
> 输入: date="2019-01-09"
>
> 输出: 9
>
> 解释: 给定日期是 2019 年的第九天。
>
> 示例2：
>
> 输入: date="2019-02-10"
>
> 输出：41

<details>
<summary style="cursor: pointer">🔑 参考：</summary>
<div>

## 分析

利用字符串分割 split 方法, 将输入的日期切分为年、月、日, 并转换为整型直接创建含有每个月份天数的列表

当为闰年时2月为29天, 将输入月份之前的每个月的天数相加,再加上输入的该月的天数,即可得到日期在当年的第多少天


## 参考代码

### Golang 代码实现

```golang
package main

import (
	"fmt"
	"strings"
	"strconv"
)

// 实例 1: "2024-03-01"
// 入口
func main() {
	data := "2024-03-01"
	fmt.Println("初始数据", data)

	solution(data)
}

// 解决方案
func solution(date string) {
	mdays := []int{
		31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31,
	}
	sp := strings.Split(date, "-")
	y, ok := strconv.Atoi(sp[0])
	if ok != nil {
		panic("month error")
	}
	m, ok := strconv.Atoi(sp[1])
	if ok != nil {
		panic("month error")
	}
	d, ok := strconv.Atoi(sp[2])
	if ok != nil {
		panic("day error")
	}

	if y % 400 == 0 || (y % 4 == 0 && y % 100 != 0) {
		mdays[1] = 29
	}
	total := d
	for i, t := range mdays[:m-1] {
		total += t
	}
	fmt.Println(date, "是", sp[0], "年的第", total, "天")
}
```

### Python 代码实现

```python
date = input("enter yyyy-mm-dd: ")
year, month, day = [int(x)for x in date.split("-")]
amount = [31,28,31,30,31,30,31,31,30,31,30,31]
if year % 400 == 0 or (year%100!=0 and year%4==0):
	# 闰年2月为29天
	amount[1] = 29
result = sum(amount[:month-1])+day
print(f"{date} 是 {year} 年的第 {result} 天")
```

```
enter yyyy-mm-dd: 2024-03-01
2024-03-01 是 2024 年的第 61 天
```

</div>
</details>
