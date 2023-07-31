# Go 每日一题

> 今日（2023-05-07）的题目如下

如何确认两个 map 是否相等？

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

map 深度相等的条件：

1. 都为 nil
2. 非空、长度相等，指向同一个 map 实体对象
3. 相应的 key 指向的 value “深度”相等

直接将使用 map1 == map2 是错误的。这种写法只能比较 map 是否为 nil。

```golang
package main

import "fmt"

func main() {
	var m map[string]int
	var n map[string]int

	fmt.Println(m == nil)
	fmt.Println(n == nil)

	// 不能通过编译
	//fmt.Println(m == n)
}
```

输出结果：

```
true
true
```

因此只能是遍历 map 的每个元素，比较元素是否都是深度相等。

答案解析来自：[https://golang.design/go-questions/map/compare/](https://golang.design/go-questions/map/compare/)

### 1楼

map是无序存储的，所以不能直接判断两个map是否相等；网上有一种比较方法，不一定是最好，但起码也是一种方案。

```golang
func CompareMap(dm1 map[string]interface{}, dm2 map[string]interface{}) bool {
    keySlice := make([]string, 0)
    data1Slice := make([]interface{}, 0)
    data2Slice := make([]interface{}, 0)
    for key, value := range dm1 {
        keySlice = append(keySlice, key)
        data1Slice = append(data1Slice, value)
    }
    for _, key := range keySlice {
        if data, ok := dm2[key]; ok {
            data2Slice = append(data2Slice, data)
        } else {
            return false
        }
    }
    data1Bytes, _ := json.Marshal(data1Slice)
    data2Bytes, _ := json.Marshal(data2Slice)
    return string(data1Bytes) == string(data2Bytes)
}
```

代码中遗漏了判断两个map是否为nil以及是否长度相等的前置条件判断，大概思路就是取出其中一个map的key放到slice里，并且按照key的slice到另外一个map中取值，如果有一个取不到，那这两个map肯定不相等,如果都取到了，那么比较两个dataslice的json编码是否相等


### 2楼

```golang
package main

import ( "fmt" "reflect" )

func main() {
    //定义两个map m := make(map[string]string) m1 := make(map[string]string)

    m["name"] = "wangshao"
    m["age"] = "15"
    m["sex"] = "man"

    m1["name"] = "wangshao"
    m1["age"] = "15"
    m1["sex"] = "man"

    //先遍历一个map拿出所有的key
    k := make([]string, 0)
    v := make([]string, 0)
    for key, value := range m {
        k = append(k, key)
        v = append(v, value)
    }

    if len(m) == len(m1) {
        for i := 0; i < len(k); i++ {
            if m1[k[i]] == v[i] {
                fmt.Println("same")
            } else {
                fmt.Println("no same")
            }
        }
    }

    //反射方式
    fmt.Println(reflect.DeepEqual(m, m1))
}
```

### 3楼

先要确定一个问题

为什么确认两个map相等

怎么算相等


### 4楼

标准答案可能不太方便，我认为，反射更加合理

```golang
package main

import( 
    "fmt"
    "relflect"
)
func main() {
    var m map[string]int
    var n map[string]int

    fmt.Println(reflect.DeepEqual(m,n))
}
```

### 18楼

凭什么先对问题做一个假设，不是string就是int？ 有没有想过为什么会有hashMightPanic？因为map不只是string、int，还有interface呢。 hash可能panic，compare也是可能panic的。结论可不只是true/false，panic不能忽略。


### 23楼

```golang
// EqualMap 比较 map 是否相等（只适用成员为标量的map）
func EqualMap(m1, m2 map[string]interface{}) bool {
    if (0 == len(m1) && 0 == len(m2)) || (nil == m1 && nil == m2) {
        return true
    }
    var keys1 []string
    for key, _ := range m1 {
        keys1 = append(keys1, key)
    }
    var keys2 []string
    for key, _ := range m2 {
        keys2 = append(keys2, key)
    }
    for _, key := range keys1 {
        if v2, ok := m2[key]; ok && m1[key] == v2 {
            continue
        }
        return false
    }
    return true
}

// EqualMapReflect 比较 map 是否深相等
func EqualMapReflect(m1, m2 map[string]interface{}) bool {
    if (0 == len(m1) && 0 == len(m2)) || (nil == m1 && nil == m2) {
        return true
    }
    return reflect.DeepEqual(m1, m2)
}
```

### 32楼

官网文档规定，slice, map, 和 function 之间不能直接比较是否相等，特例是 nil。具体参照：[https://go.dev/ref/spec#Comparison_operators](https://go.dev/ref/spec#Comparison_operators)

Go 中 map 的 value 类型可以是任意类型的，这就把上面除了『reflect.DeepEqual』之外的其他方法几乎都毙掉了，因为要判断value 类型是 slice、map 或 function 的话该如何处理。要自己实现的话，其实参照 reflect.DeepEqual 的实现来比较好。


</div>
</details>
