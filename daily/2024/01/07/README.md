# 2024-01-07 练习题

> 来源 : 2023-12-30 [Python 每日一练-LeetCode-数组练习-是否组成最多糖果数](https://www.bilibili.com/video/BV16K411b7Cr/)

## 题目说明

- 给你一个数组 candies 和一个整数 extraCandies，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。
- 对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有最多的糖果。注意，允许有多个孩子同时拥有最多的糖果数目。

示例:

> 输入: candies = [4, 2, 1, 1, 2], extraCandies = 1
>
> 输出: [true, false, false, false, false]
>
> 解释: 只有 1 个额外糖果，所以不管额外糖果给谁，只有孩子 1 可以成为拥有糖果最多的孩子。

> 输入: candies = [12, 1, 12], extraCandies = 10
>
> 输出: [true, false, true]

## 说明

> 输入: candies = [2, 3, 5, 1, 3], extraCandies = 3
>
> 输出: [true, true, true, false, true]
>
> 解释:

- 孩子 1 有 2 个糖果，如果他得到所有额外的糖果(3 个)，那么他总共有 5 个糖果，他将成为拥有最多糖果的孩子。
  > - 孩子 2 有 3 个糖果，如果他得到至少 2 个额外糖果，那么他将成为拥有最多糖果的孩子。
  > - 孩子 3 有 5 个糖果，他已经是拥有最多糖果的孩子。
  > - 孩子 4 有 1 个糖果，即使他得到所有额外的糖果，他也只有 4 个糖果，无法成为拥有糖果最多的孩子。
  >   孩子 5 有 3 个糖果，如果他得到至少 2 个额外糖果，那么他将成为拥有最多糖果的孩子。

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
	"sort"
)

// 定义数据结构
type CandyData struct {
	Candies      []int
	ExtraCandies int
}

// 入口
func main() {
	// 所有示例测试
	// all_candies := []CandyData{
	// 	{Candies: []int{4,2,1,1,2}, ExtraCandies: 1},
	// 	{Candies: []int{12,1,12}, ExtraCandies: 10},
	// 	{Candies: []int{2,3,5,1,3}, ExtraCandies: 3},
	// }
	// for _, data := range all_candies {
	// 	fmt.Println("初始数据", data)
	// 	solution(data.Candies, data.ExtraCandies)
	// }

	data := CandyData{Candies:[]int{4,2,1,1,2}, ExtraCandies:1}
	fmt.Println("初始数据", data)

	// max 函数需要 golang v 1.21 版本
	// fmt.Printf("最大糖果数: %v\n", max(data.Candies))
	solution(data.Candies, data.ExtraCandies)
}

// 解决方案
func solution(candies []int, extraCandies int) {
	// 深拷贝
	bak := make([]int, len(candies))
	copy(bak, candies)
	// 排序
	sort.Ints(bak)
	debug_log("最多的糖果数量:", bak[len(bak)-1])
	for inx, candy := range candies {
		fmt.Printf("%d. %d + %d = %d", inx, candy, extraCandies, candy+extraCandies)
		if candy+extraCandies >= bak[len(bak)-1] {
			fmt.Println(" 分配后糖果数量是最多 (true)")
		} else {
			fmt.Println(" 分配后糖果数量不是最多 (false)")
		}
	}
}
```

### Python 代码实现

#### 方法一: 详细输出

```python
# 详细输出
candies = [4,2,1,1,2]
extraCandies = 1
print(f"candies列表中最多糖果的数量是{max(candies)}")
s=[]
for candy in candies:
    ls.append((f"{candy}+{extraCandies}={candy+extraCandies}", candy+extraCandies >= max(candies)))
print(ls)
```

#### 方法二: 简化输出

```python
# 简化输出
candies = [4,2,1,1,2]
extraCandies = 1
ls = [candy+extraCandies>=max(candies) for candy in candies]
print(ls)
```

</div>
</details>
