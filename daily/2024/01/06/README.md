# 2024-01-05 练习题

> 来源 : 2023-12-31 [Python 每日一练-LeetCode-数组练习-最大乘积](https://www.bilibili.com/video/BV1Bg4y117MF/)

## 题目说明

> 给你一个整数数组 nums，请你选择数组的两个不同下标 i 和 j，使 `(nums[i]-1)*(nums[j]-1)` 取得最大值。
>
> 请你计算并返回该式的最大值。


## 说明

> 输入: nums = [3, 4, 5, 2]
>
> 输出: 12
>
>解释: 如果选择下标 i=1 和 j=2 (下标从0开始)，则可以获得最大值，`(nums[1]-1)*(nums[2]-1) = (4-1)*(5-1) = 3*4 = 12`。

> 输入: nums = [1, 5, 4, 5]
>
> 输出: 16
>
>解释: 选择下标 i=1 和 j=3 (下标从0开始)，则可以获得最大值 `(5-1)*(5-1) = 16`。

> 输入: nums = [3, 7]
>
> 输出: 12


<details>
<summary style="cursor: pointer">🔑 参考：</summary>
<div>

## 分析

找到给定数列中的最大值和次大值，然后进行运算即可。


## 参考代码

### Golang 代码实现

```golang
import (
	"fmt"
	"sort"
)

func main() {
	// 所有示例测试
	all_nums := [][]int{
		{3, 4, 5, 2},
		{1, 5, 4, 5},
		{3, 7},
	}
	for _, nums := range all_nums {
		fmt.Println("初始数据", nums)
		solution(nums)
		solution_b(nums)
	}
}

// 解决方案 1
func solution(nums []int) {
	sort.Ints(nums)
	x, y := nums[len(nums)-1], nums[len(nums)-2]
	fmt.Printf("方法1: 最大两个数各自减1后乘积为: (%d - 1) * (%d - 1) = %d * %d = %d\n", x, y, x-1, y-1, (x-1)*(y-1))
}

// 解决方案 2
func solution_b(nums []int) {
	if len(nums) < 2 {
		fmt.Println("数组长度必须大于等于2")
		return
	}
	x, y := nums[0], nums[1]
	if x < y {
		x, y = y, x
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > x {
			x, y = nums[i], x
		} else if nums[i] < y {
			y = nums[i]
		}
	}
	fmt.Printf("方法2: 最大两个数各自减1后乘积为: (%d - 1) * (%d - 1) = %d * %d = %d\n", x, y, x-1, y-1, (x-1)*(y-1))
}
```

### Python 代码实现

#### 方法一: 因为数组中全为正整数，只需要输出乘积

```python
nums = [3, 4, 5, 2]
nums.sort()
print(f"最大两个数各自减1后乘积为："
	f"{nums[-1]-1}*{nums[-2]-1}={(nums[-1]-1)*(nums[-2]-1)}")
```

#### 方法二: 一次遍历维护最大和次大值

> - 因为我们只需要得到数组中两个最大的元素,
> - 我们可以在从左到右遍历的过程中维护两个变量 a, b 来表示遍历过程中的最大和次大元素,
> - 那么一次遍历就可以得到数组中两个最大的元素.

```python
nums = [3, 4, 5, 2]
maxB, secB = nums[0], nums[1]
if maxB < secB:
	maxB, secB = secB, maxB

for i in range(2, len(nums)):
	num = nums[i]
	if num > maxB:
		maxB, secB = num, maxB
	elif num > secB:
		secB = num

print(f"最大数为{maxB}, 次大数为{secB},"
	f"各自减1后乘积为{(maxB-1)*(secB-1)}")
```

</div>
</details>
