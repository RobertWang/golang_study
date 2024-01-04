# 2024-01-02 练习题

> 来源 : [Python 每日一练-LeetCode-数组练习-石头最后的重量](https://www.bilibili.com/video/BV1kw411x7JT/)

## 题目说明

> - 有一堆石头，每块石头的重量都是正整数。
> - 每一回合，从中选出两块最重的石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
> - 如果 x == y，那么两块石头都会被完全粉碎；
> - 如果 x !=y，那么重量为 x 的石头将会完全粉碎， 而重量为 y 的石头新重量为 y-x。
> - 最后，最多只会剩下一块石头。返回此石头的重量。 如果没有石头剩下，就返回 0。

## 示例说明

- 输入：[2, 7, 4, 1, 8, 1]
- 输出：1
- 解释：
  - 先选出 7 和 8，得到 1，所以数组转换为 [2, 4, 1, 1, 1]，
  - 再选出 2 和 4，得到 2，所以数组转换为 [2, 1, 1, 1]，
  - 接着是 2 和 1，得到 1，所以数组转换为 [1, 1, 1]，
  - 最后选出 1 和 1，得到 0，最终数组转换为 [1] ，这就是最后剩下那块石头的重量。

<details>
<summary style="cursor: pointer">🔑 参考：</summary>
<div>

## 分析

- 将所有石头的重量放入最大堆中。每次依次从队列中取出最重的两块石头 a 和 b，必有 a ≥ b。
- 如果 a > b，则将新石头 a-b 放回到最大堆中；
- 如果 a = b，两块石头完全被粉碎，因此不会产生新的石头。
- 重复上述操作，直到剩下的石头少于 2 块。
- 最终可能剩下 1 块石头，该石头的重量即为最大堆中剩下的元素，返回该元素；
- 也可能没有石头剩下，此时最大堆为空，返回 0。

## 参考代码

### Golang 代码实现

```golang
import "fmt"
func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println("初始数据", stones)

	solution(stones)
}

// 解决方案
func solution(stones []int) {
	sort.Sort(sort.IntSlice(stones))

	times := 0
	for len(stones) > 1 {
		times = times + 1
		// 判断最大的数字 x 与 第2大的数字 y 是否相同
		if stones[len(stones)-1] > stones[len(stones)-2] {
			// 不相同，则第2大的数字变更为 x - y
			stones[len(stones)-2] = stones[len(stones)-1] - stones[len(stones)-2]
			stones = stones[:len(stones)-1]
			// 重新排序
			sort.Sort(sort.IntSlice(stones))
		} else {
			// 如相同，则 2 个数字全都清除
			stones = stones[:len(stones)-2]
		}
	}

	if len(stones) == 0 {
		fmt.Println("最终没有石头剩下")
	} else {
		fmt.Println("最后剩下的石头重量为", stones[0])
	}
}
```

### Python 代码实现

```python
stones = [2, 7, 1, 8, 1]
st = sorted(stones)
while len(st)>1:
	if st[-1]-st[-2] == 0:
		st=st[:-2]
	else:
		st = sorted(st[:-2]+[st[-1]-st[-2]])
if st==[]:
	print("石头全部被粉碎")
else:
	print(f"剩下的石头为{st[0]}")
```

</div>
</details>
