# Go每日一题

> 今日（2023-02-01）的题目如下

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


### 14楼

- A. 单引号里面应该是单个字符 类型是rune 类型的
- C. 就是在A的基础上，两个类型完全不一样没法直接相加


</div>
</details>