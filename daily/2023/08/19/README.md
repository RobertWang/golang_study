# Go 每日一题

> 今日（2023-08-19）的题目如下

下面这段代码输出什么，说明原因。

```golang
func main() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for key,val := range slice {
		m[key] = &val
	}

	for k,v := range m {
		fmt.Println(k,"->",*v)
	}
}
```

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

```
0 -> 3
1 -> 3
2 -> 3
3 -> 3
```

解析：这是新手常会犯的错误写法，for range 循环的时候会**创建每个元素的副本，而不是元素的引用**，所以 m[key] = &val 取的都是变量 val 的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为 3，所有输出都是 3.

正确的写法：

```golang
func main() {

	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for key,val := range slice {
		value := val
		m[key] = &value
	}

	for k,v := range m {
		fmt.Println(k,"===>",*v)
	}
}
```

扩展题目

```golang
type Test struct {
	name string
}

func (this *Test) Point(){
	fmt.Println(this.name)
}

func main() {

	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}

	for _,t := range ts {
		//fmt.Println(reflect.TypeOf(t))
		defer t.Point()
	}

}
```

参考：[https://blog.csdn.net/idwtwt/article/details/87378419](https://blog.csdn.net/idwtwt/article/details/87378419)

---

### 2 楼

楼上正解. 因为循环体内的 val 为同一个值(处于同一个地址,只是循环时不断改变值). 因此, `&val` 是固定的. 循环结束时 `val == 3`, 因此最后所有的 `*v == 3`

</div>
</details>
