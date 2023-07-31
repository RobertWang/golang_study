# Go每日一题

> 今日（2023-02-15）的题目如下

关于 channel，下面语法正确的是：

- A. var ch chan int
- B. ch := make(chan int)
- C. <- ch
- D. ch <-



<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：ABC。

A、B 都是声明 channel；C 读取 channel；写 channel 是必须带上值，所以 D 错误。

### 14楼

个人觉得严谨点应该只有A是对的。D是错误的不用说，BC都是要在func内才是正确的，如果在func外就是不合法的。

### 19楼

`wg.Add(1)` 上移一行。够执行太快，可能来不及执行，造成后面的 `wg.Wait()` 拦不住。

</div>
</details>
