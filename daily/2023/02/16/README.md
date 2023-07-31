# Go每日一题

> 今日（2023-02-16）的题目如下

下面这段代码输出什么？

```golang
type person struct {  
    name string
}

func main() {  
    var m map[person]int
    p := person{"mike"}
    fmt.Println(m[p])
}
```

- A. 0
- B. 1
- C. Compilation error



<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：A。

m 是一个 map，值是 nil。从 nil map 中取值不会报错，而是返回相应的零值，这里值是 int 类型，因此返回 0。

### 19楼

`map[p]` 在经过 编译后，走的是 `runtime.mapaccess1` 的逻辑；而看到 `mapaccess1` 函数，对于 `hmap` 是 `nil` 的情况是直接返回空值；源代码如下：

```golang
func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
    // ...

    if h == nil || h.count == 0 {// h 就是map指向的地址，因为题目中map还没有申请分配内存空间，所以h是nil
        if t.hashMightPanic() {
            t.hasher(key, 0) // see issue 23734
        }
        return unsafe.Pointer(&zeroVal[0])
    }

    // ...
}
```

</div>
</details>
