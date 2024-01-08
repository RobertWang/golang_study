# 2024-01-08 练习题

> 来源 : 2023-12-29 [Python 每日一练-LeetCode-数组练习-提莫攻击](https://www.bilibili.com/video/BV14c411m74B/)

## 题目说明

- 在《英雄联盟》的世界中，有一个叫 “提莫” 的英雄。他的攻击可以让敌方英雄艾希 (编者注: 寒冰射手) 进入中毒状态。
- 当提莫攻击艾希，艾希的中毒状态正好特续 duration 秒。
- 正式地讲，提莫在 t 发起攻击意味着艾希在时间区间 [t, t+duration-1] (含 t 和 t+duration-1) 处于中毒状态。
- 如果提莫在中毒影响结束前再次攻击，中毒状态计时器将会重置，在新的攻击之后，中毒影响将会在 duration 秒后结束。
- 给你一个非递减的整数数组 timeSeries，其中 timeSeries[i] 表示提莫在 timeSeries[i] 秒时对艾希发起攻击，以及一个表示中毒持续时间的整数 duration。
- 返回艾希处于中毒状态的总秒数。

## 示例说明

> 输入: timeSeries = [1, 4], duration = 2
>
> 输出: 4
>
> 解释: 提莫攻击对艾希的影响如下：
>
> - 第 1 秒，提莫攻击艾希并使其立即中毒。中毒状态会维持 2 秒，即第 1 秒和第 2 秒。
> - 第 4 秒，提莫再次攻击艾希，艾希中毒状态又持续 2 秒，即第 4 秒和第 5 秒。艾希在第 1、2、4、5 秒处于中毒状态，所以总中毒秒数是 4。

> 输入: timeSeries = [1, 2], duration = 2
>
> 输出: 3
>
> 解释: 提莫攻击对艾希的影响如下：
>
> - 第 1 秒，提莫攻击艾希并使其立即中毒。中毒状态会维持 2 秒，即第 1 秒和第 2 秒。 > - 第 2 秒，提莫再次攻击艾希，并重置中毒计时器，艾希中毒状态需要持续 2 秒，即第 2 秒和第 3 秒。
>   艾希在第 1、2、3 秒处于中毒状态，所以总中毒秒数是 3。

<details>
<summary style="cursor: pointer">🔑 参考：</summary>
<div>

## 分析

- 我们只需要对数组进行一次扫描就可以计算出总的中毒持续时间。
- 我们记录艾希恢复为未中毒的起始时间 expired，设艾希遭遇第 i 次的攻击的时间为 timeSeries[i]。
- 当艾希遭遇第 i 次攻击时: 如果当前他正处于未中毒状态，则此时他的中毒持续时间应增加 duration，同时更新本次中毒结束时间 expired 等于 timeSeries[i]+duration；
- 如果当前他正处于中毒状态，由于中毒状态不可叠加，我们知道上次中毒后结束时间为 expired，本次中毒后结束时间为 timeSeries[i]+duration，因此本次中毒增加的持续中毒时间为 timeSeries[i]+duration-expired；
- 我们将每次中毒后增加的特续中毒时间相加即为总的持续中毒时间。

## 参考代码

### Golang 代码实现

```golang
package main

import (
	"fmt"
)

// 定义数据结构
type Data struct {
	TimeSeries []int
	Duration   int
}

// 入口
func main() {
	data := Data{TimeSeries: []int{1, 2, 4}, Duration: 2}
	fmt.Println("初始数据", data)

	solution(data.TimeSeries, data.Duration)
}

// 解决方案
func solution(timeSeries []int, duration int) {
	total, expired := 0, 0
	for _, point := range timeSeries {
		if point > expired {
			total += duration
		} else {
			total += point + duration - expired
		}
		expired = point + duration
	}
	fmt.Println("提莫攻击所致中毒的总时间为", total)
}
```

### Python 代码实现

```python
timeSeries = [1,2,4]
duration = 2
time, expired = 0,0
for i in range(len(timeSeries)):
    if timeSeries[i]>expired:
        time += duration
    else:
        time += timeSeries[i]+duration-expired
    expired = timeSeries[i]+duration
print(f"提莫攻击所致中毒时间总长为{time}")
```

</div>
</details>
