# Go 每日一题

> 今日（2023-05-10）的题目如下

new() 与 make() 的区别

<details>
<summary>答案解析：</summary>
<div>

new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。

new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。

make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel。

---

### 1 楼

1. new 返回对象的指针
2. make 只能初始化 channel map slice

### 7 楼

make 只能针对map、slice、chanel初始化，初始化出来。这4种var完后必须make才能使用。 new 针对其他变量，例如数组。

### 13 楼

【1】 初始化对象不同，make 主要是初始化 切片、通道、map等。【2】new 初始化出来只是一个指针，不分配内存。make 直接分配内存。

</div>
</details>
