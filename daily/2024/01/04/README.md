# 2024-01-04 练习题

[Python 每日一练-LeetCode-数组练习-寻找终点站](https://www.bilibili.com/video/BV13e411B78X/)

## 题目说明

- 给你一份旅游线路图，该线路图中的旅行线路用数组 paths 表示，
- 其中 paths[i] = [cityAi, cityBi] 表示该线路将会从 cityAi 直接前往 cityBi。
- 请你找出这次旅行的终点站，即没有任何可以通往其他城市的线路的城市。
- 题目数据保证线路图会形成一条不存在循环的线路，因此恰有一个旅行终点站。

## 说明

> 输入: paths = [["London", "New York"], ["New York", "Lima"], ["Lima", "Sao Paulo"]]
>
> 输出: "Sao Paulo"
>
> 解释: 从 "London" 出发, 最后抵达终点站 "Sao Paulo". 本次旅行的路线是 "London" -> "New York" -> "Lima" -> "Sao Paulo"。

> 输入: paths = [["B", "C"], ["D", "B"], ["C", "A"]]
>
> 输出: "A"
>
> 解释: 所有可能的线路是:
>
> - "D" -> "B" -> "C" -> "A".
> - "B" -> "C" -> "A".
> - "C" -> "A".
> - "A".
>   显然, 旅行终点站是"A".

> 输入: paths = [["A","Z"]]
>
> 输出: "Z"

<details>
<summary style="cursor: pointer">🔑 参考：</summary>
<div>

## 分析

- 根据终点站的定义，终点站不会出现在 cityAi 中，因为存在从 cityAi 出发的线路，所以终点站只会出现在 cityBi 中。
- 据此，我们可以遍历 cityBi，返回不在 cityAi 中的城市，即为答案。
- 代码实现时，可以先将所有 cityAi 存于一哈希表中，然后遍历 cityBi 并查询 cityBi 是否在哈希表中。

## 参考代码

### Golang 代码实现

```golang
import "fmt"

// 入口
func main() {
	data := [][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	}
	fmt.Println("初始数据", data)

	solution(data)
}

// 解决方案
func solution(data [][]string) {
	starts := make(map[string]int)
	for inx, points := range data {
		starts[points[0]] = inx
	}
	exists := true
	for _, points := range data {
		_, exists = starts[points[1]]
		if !exists {
			fmt.Println("终归点城市是", points[1])
		}
	}
	if exists {
		fmt.Println("当前城市路径规划没有终归")
	}
}
```

### Python 代码实现

```python
paths = [["London", "New York"], ["New York", "Lima"], ["Lima", "Sao Paulo"]]
finalstation = set()

for path in paths:
	finalstation.add(path[0])

for path in paths:
	if path[1] not in finalstation:
		print(f"the final station is {path[1]}")

# 优化
# finalstation = {path[0] for path in paths}
# print(next(path[1] for path in paths if path[1] not in finalstation))
```

</div>
</details>
