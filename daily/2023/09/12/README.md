# Go 每日一题

> 今日（2023-09-12）的题目如下

关于字符串连接，下面语法正确的是？

- A. str := 'abc' + '123'
- B. str := "abc" + "123"
- C. str := '123' + "abc"
- D. fmt.Sprintf("abc%d", 123)

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案：BD

参考解析：考的知识点是字符串连接。除了以上两种连接方式，还有 strings.Join()、buffer.WriteString() 等。

</div>
</details>
