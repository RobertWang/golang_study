# 2024-01-18 练习题

> 来源 : 2024-01-18 [Python 每日一练-LeetCode-字符串练习-动态密码](https://www.bilibili.com/video/BV14g4y1m7Yi/)

## 题目说明

某公司门禁密码使用动态口令技术。初始密码为字符串 password，密码更新均遵循以下步骤:

- 设定一个正整数目标值 target
- 将 password 前 target 个字符按原顺序移动至字符串末尾
- 返回更新后的密码字符串

## 示例说明

> 示例 1:
>
> 输入: password="s3cur1tyC0d3",target=4
>
> 输出: "r1tyC0d3s3cu"
>
> 示例 2:
>
> 输入: password="Irloseumgh",target=6
>
> 输出: "umghlrlose"

<details>
<summary style="cursor: pointer">🔑 参考：</summary>
<div>

## 分析

输入原始密码和 target

将原始密码字符串的前 target 个字符通过切片移到最后, 原始字符的从 target 索引往后的切片移到开始即可.

## 参考代码

### Golang 代码实现

```golang
package main

import (
	"fmt"
)

type PasswordData struct {
	Password string
	Target   int
}

func main() {
	data := PasswordData{Password:"s3cur1tyC0d3", Target:4}
	fmt.Println("初始数据", data)

	solution(data.Password, data.Target)
}

// 解决方案
func solution(password string, target int) {
	result := password[target:] + password[:target]
	fmt.Println(result)
}
```

### Python 代码实现

```python
password = input("enter origin password:")
target = int(input("enter target number:")
print(f"new password:{password[target:]+password[:target]}")
```

```
enter origin password: s3curitycod3
enter target number: 4
new password: r1tycod3s3cu
```

</div>
</details>
