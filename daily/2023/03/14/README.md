# Go 每日一题

> 今日（2023-03-14）的题目如下

下面代码中 A B 两处应该怎么修改才能顺利编译？

```golang
func main() {
	var m map[string]int        //A
	m["a"] = 1
	if v := m["b"]; v != nil {  //B
		fmt.Println(v)
	}
}
```


<details>
<summary>答案解析：</summary>
<div>

参考答案及解析：

```golang
func main() {
	m := make(map[string]int)
	m["a"] = 1
	if v,ok := m["b"]; ok {
		fmt.Println(v)
	}
}
```

在 A 处只声明了map m ,并没有分配内存空间，不能直接赋值，需要使用 make()，都提倡使用 make() 或者字面量的方式直接初始化 map。

B 处，v,k := m["b"] 当 key 为 b 的元素不存在的时候，v 会返回值类型对应的零值，k 返回 false。


</div>
</details>
