# Go 每日一题

> 今日（2023-05-04）的题目如下

Go 的 map 可以边遍历边删除吗？

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

map 并不是一个线程安全的数据结构。同时读写一个 map 是未定义的行为，如果被检测到，会直接 panic。

上面说的是发生在多个协程同时读写同一个 map 的情况下。 如果在同一个协程内边遍历边删除，并不会检测到同时读写，理论上是可以这样做的。但是，遍历的结果就可能不会是相同的了，有可能结果遍历结果集中包含了删除的 key，也有可能不包含，这取决于删除 key 的时间：是在遍历到 key 所在的 bucket 时刻前或者后。

一般而言，这可以通过读写锁来解决：sync.RWMutex。

读之前调用 RLock() 函数，读完之后调用 RUnlock() 函数解锁；写之前调用 Lock() 函数，写完之后，调用 Unlock() 解锁。

另外，sync.Map 是线程安全的 map，也可以使用。

参考答案来自：[https://golang.design/go-questions/map/delete-on-range/](https://golang.design/go-questions/map/delete-on-range/)

### 1楼

清空map， :smile:

```golang
for k := range m {
    delete(m, k)
}
```

### 12楼

```golang
// Go1.11版本以上这种清空map方法有效
for k := range mapdemo {
    delete(mapdemo , k)
}
```

### 15楼

@Dessert</a> 我记得有一期的每日一题研究过range遍历，map的遍历是实时的，在遍历第1个元素时删除第2个元素，那么后续就不会遍历第2个元素。遍历第1个元素时删除第1个元素，后续更不会再出现第1个元素了。

[https://go.dev/doc/effective_go#for](https://go.dev/doc/effective_go#for) ， 这个官方例子也展示了可以在遍历的时候删除。

> If you only need the first item in the range(the key or index),drop the second:
> ```golang
> for key :range m {
>     if key.expired(){
>         delete(m,key)
>     }
> }
> ```

[https://go.dev/ref/spec#For_statements](https://go.dev/ref/spec#For_statements) ， 同时官方的range迭代也有说着遍历时删除和新增的情况

![](https://static.golangjob.cn/220923/7a7bd664822e26524057441a9d68c8a0.png)

我感觉清空map还是直接用`m=make(map[string]string)`生成新对象，让GC清理旧map好点，因为map的delete并不会真的删除里面元素，貌似只是标记被删除，这个比较底层没深入研究，这时还是会占用一些内存吧。

### 17楼

多个协程同时读写同一个 map，会得到如下的panic噢

```
fatal error: concurrent map iteration and map write
```

多协程下可以使用Go1.9版本引入的sync.Map类型来替换map

> 💡Tips：多协程读map是没问题的，但是写不行

### 20楼

> 回复15楼

map 是引用，所以虽然 walkrange 生成了一个 ha，但是修改 map，ha 也会收到影响。delete 只是把 tophash置为 emptyOne 或者 emptyRest，本质上 key 和 value 数组部分占用的内存还在，所以并不是真的删除，clear 会进行分配

### 23楼

> 回复15楼

那就是做好直接通过赋值新map给原来的map ‘m’来清空是么， 剩下的交给GC


### 24楼

23楼 @Dessert 是的，因为map的底层是hash之，存数据越多的map后续存取应该会慢点，map重复使用个人感觉后续存值和取值的性能都会受到影响。除非像#20楼说的触发了map的clear机制，里面数据真的被清理掉。

</div>
</details>
