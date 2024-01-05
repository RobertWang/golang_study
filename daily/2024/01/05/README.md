# 2024-01-05 练习题

> 来源 : 2024-01-01 [Python 每日一练-LeetCode-数组练习-寻找小镇法官](https://www.bilibili.com/video/BV1Ya4y1r7UN/)

## 题目说明

>     小镇里有 n 个人，按从 1 到 n 的顺序编号。传言称， 这些人中有一个暗地里是小镇法官。
>
>     如果小镇法官真的存在，那么：
>
>     - 小镇法官不会信任任何人。
>     - 每个人（除了小镇法官）都信任这位小镇法官。
>     - 只有一个人同时满足属性 1 和属性 2。
>     - 给你一个数组 trust, 其中trust[i] = [ai，bi]表示编号为 ai 的人信任编号为 bi 的人。
>     - 如果小镇法官存在并且可以确定他的身份，请返回该法官的编号；否则，返回-1。

## 说明

> 输入: n=2, trust=[[1,2]]
>
> 输出: 2

> 输入: n=3, trust=[[1,3],[2,3]]
>
> 输出: 3

> 输入: n=3, trust=[[1,3],[2,3],[3,1]]
>
> 输出: -1

> 输入: n=4, trust=[[1,3],[1,4],[2,3],[2,4],[4,3]]
>
> 输出: 3

> 本题需要用到有向图中节点的入度和出度的概念。在有向图中，一个节点的入度是指向该节点的边的数量；而一个节点的出度是从该节点出发的边的数量。

<details>
<summary style="cursor: pointer">🔑 参考：</summary>
<div>

## 分析

- 每个人是图的节点，trust 的元素 trust[i] 是图的有向边，从 trust[i][O]指向 trust[i][1]。
- 我们可以遍历 trust, 统计每个节点的入度和出度, 存储在 inDegrees 和 outDegrees 中。
- 根据题意，在法官存在的情况下，法官不相信任何人，每个人（除了法官外）都信任法官， 且只有一名法官。
- 因此法官这个节点的入度是 n-1，出度是 0。
- 我们可以遍历每个节点的入度和出度，如果找到一个符合条件的节点，由于题目保证只有一个法官，我们可以直接返回结果；如果不存在符合条件的点，则返回一 1。

## 参考代码

### Golang 代码实现

```golang
import "fmt"
func main() {
	persons := [][]int{{1,3},{1,4},{2,3},{2,4},{4,3}}
	fmt.Println("初始数据", persons)

	solution(persons)
}

// 解决方案
func solution(persons [][]int) {
	judge := -1
	var in, out, all map[int]int
	in = make(map[int]int)
	out = make(map[int]int)
	all = make(map[int]int)
	for _, rel := range persons {
		if _, ok := all[rel[0]]; !ok {
			all[rel[0]] = 0
		}
		if _, ok := all[rel[1]]; !ok {
			all[rel[1]] = 0
		}
		if _, ok := in[rel[0]]; ok {
			out[rel[0]]++
		} else {
			out[rel[0]] = 0
		}
		if _, ok := in[rel[1]]; ok {
			in[rel[1]]++
		} else {
			in[rel[1]] = 1
		}
	}
	person_count := len(all)
	for person, _ := range all {
		in_count, out_count := 0, 0
		out_count, _ = out[person]
		in_count, _ = in[person]
		if out_count == 0 && in_count == person_count-1 {
			judge = person
			break
		}
	}
	if judge == -1 {
		fmt.Println("结果: 小镇没有法官")
	} else {
		fmt.Println("结果: 法官是", judge)
	}
}
```

### Python 代码实现

#### 官方题解

官方解答所需函数:

1.  导入 `collections.Counter` 模块

    ```python
    import collections from Counter
    ```

2.  ` Counter()` 函数: 统计 python 列表对象中每个元素出现的次数, 返回一个 `Collections.Counter` 类型数据, `{元素: 出现次数}`

3.  `next()` 函数: 返回迭代器的下一个项目.

    >     next(iterable[,default])
    >
    >     参数说明:
    >
    >     iterable : 可迭代对象
    >
    >     default : 可选, 用于设置在没有下一个元素时返回该默认值, 如果不设置, 又没有下一元素则会触发 Stoplteration 异常.

```python
from collections import Counter
trust = [[1,3][14],[2,3][2.4].[4,3]]
n = 4

inDegrees = Counter(y for x,y in trust)
outDegrees = Counter(x for x,y in trust)
print(next((i for i in range(1, n+1) if inDegrees[i]==n-1 and outDegrees[i]==0), -1))
```

#### 方法二

```python
from collections import Counter
trust = [[1,3],[1,4],[2,3],[2,4],[4,3]]
n = 4

inDegrees = Counter(y for x,y in trust)
outDegrees = Counter(x for x,y in trust)

judge = 0
for i in range(1, n+1):
	if inDegrees[i]==n-1 and outDegrees[i]==0:
		judge i
		break
	else:
		judge = -1
print(judge)
```

</div>
</details>
