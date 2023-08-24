# Go 每日一题

> 今日（2023-08-24）的题目如下

下面这段代码有什么缺陷：

```golang
func sum(x, y int)(total int, error) {
	return x+y, nil
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

答案：第二个返回值没有命名。

解析：

在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。如果有多个返回值必须加上括号()；如果只有一个返回值且命名也必须加上括号()。这里的第一个返回值有命名 total，第二个没有命名，所以错误。

---

### 3 楼

定義回傳值的部份，int 參數已給予變數名稱 (total)，error 參數也要給予變數名稱。 或者 int 參數的變數名稱也拿掉。

因為裡面實際回傳是使用 return x+y, nil，並未使用回傳變數名稱。 所以應該拿掉變數名稱比較恰當。

```golang
func sum(x, y int)(int, error) {
    return x+y, nil
}
```

### 17 楼

多个返回值，有一个具名了，剩下的都要具名

### 37 楼

> 原始代码可能会溢出，改为下面的代码更合理

```golang
func add[T ~int | ~int8 | ~int16 | ~int32 | ~int64](x, y T) T {
    r := x + y
    if (x^r)&(y^r) < 0 {
        panic("overflow")
    }
    return r
}

func subtract[T ~int | ~int8 | ~int16 | ~int32 | ~int64](x, y T) T {
    r := x - y
    if (x^y)&(x^r) < 0 {
        panic("overflow")
    }
    return r
}
```

### 38 楼

需要全都命名 另外 error 如果一直时 nil 可以取消这个返回值 精简函数

</div>
</details>
