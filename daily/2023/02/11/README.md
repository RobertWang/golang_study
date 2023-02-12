# Go每日一题

> 今日（2023-02-11）的题目如下

下面这段代码输出什么以及原因？

```golang
func hello() []string {  
    return nil
}

func main() {  
    h := hello
    if h == nil {
        fmt.Println("nil")
    } else {
        fmt.Println("not nil")
    }
}
```

- A. nil
- B. not nil
- C. compilation error


<details>
<summary>答案解析：</summary>
<div>

答案及解析：B。

这道题目里面，是将函数 hello 赋值给变量 h，而不是函数的返回值（即不是进行函数调用），所以输出 not nil。注意 Go 中函数是一等公民。

---

### 1楼

如果是

```golang
h := hello()
```

答案就应该是 A

函数 hello 复制给 h，h 得到的不是 hello 函数的返回值，所以答案是 B


### 7楼

hello() 返回值类型是 `[]string`，是 slice 类型。golang 中两个 slice 类型是不可比较的，但是 slice 可以跟 nil 比较。`[]string` 的零值是 nil

```golang
var a []string
fmt.Printf("%T \t %v \t %v\n", a, a, a == nil) // []string         []      true

var b = make([]string, 0)
fmt.Printf("%T \t %v \t %v\n", b, b, b == nil) // []string         []      false
fmt.Printf("%v\n", a == b)  // Invalid operation: a == b (the operator == is not defined on []string)
```

### 10楼

h := hello 是函数指针故不为 nil

---

### 补充知识

**如何理解 Go 中函数是一等公民**

- [Go 语言中的一等公民：看似普通的函数，凭什么？](https://blog.csdn.net/eddycjy/article/details/114985996)
- [一文搞懂golang函数高级用法:匿名、闭包及高阶函数](https://cloud.tencent.com/developer/article/1908081)

</div>
</details>
