# Go每日一题

> 今日（2023-02-10）的题目如下

关于 init 函数，下面说法正确的是：

- A. 一个包中，可以包含多个 init 函数；
- B. 程序运行时，先执行依赖包的 init 函数，再执行 main 包内的 init 函数；
- C. main 包中，不能有 init 函数；
- D. init 函数可以被其他函数调用；

<details>
<summary style="cursor: pointer">🔑 答案解析：</summary>
<div>

参考答案及解析：AB。

关于 init() 函数有几个需要注意的地方：

- init() 函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等;
- 一个包可以出~~线~~现多个 init() 函数,一个源文件也可以包含多个 init() 函数；
- 同一个包中多个 init() 函数的执行顺序没有明确定义，但是不同包的init函数是根据包导入的依赖关系决定的;
- init() 函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译错误;
- 一个包被引用多次，如 A import B, C import B, A import C，B 被引用多次，但 B 包只会初始化一次；
- 引入包，不可出现死循坏。即 A import B, B import A，这种情况编译失败；


### 4楼

学到的知识点： 一个包中、一个源文件中都可以出现多个init函数。

补充go文件的初始化顺序：

1. 引入的包
2. 当前包中的常量变量
3. 当前包的init
4. main函数

### 6楼

B 的说法不严谨，编译的时候，不会执行 init，运行的时候才执行


</div>
</details>
