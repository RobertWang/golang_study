# Go 每日一题

> 今日（2023-06-06）的题目如下

以下代码有什么问题？

```golang
package main

import (
	"sync"
)

const N = 10

var wg = &sync.WaitGroup{}

func main() {
	for i := 0; i < N; i++ {
		go func(i int) {
			wg.Add(1)
			println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
```


<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

输出结果不唯一，代码存在风险, 所有 go 语句未必都能执行到。

这是使用 WaitGroup 经常犯下的错误！请各位同学多次运行就会发现输出都会不同甚至又出现报错的问题。 这是因为 go 执行太快了，导致 wg.Add(1) 还没有执行 main 函数就执行完毕了。wg.Add 的位置放错了。

改为下面代码试试：

```golang
package main

import (
	"sync"
)

const N = 10

var wg = &sync.WaitGroup{}

func main() {

    for i:= 0; i< N; i++ {
        wg.Add(1)
        go func(i int) {
            println(i)
            defer wg.Done()
        }(i)
    }

    wg.Wait()
}
```



</div>
</details>
