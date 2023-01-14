// main.go
package main

import (
	"./subpack" //这里是目录相对路径,不是包名
	"fmt"
)

func main() {
	fmt.Println(subpack.Sub()) //这里的requests才是./requests目录中的package名称
	function()                 //因为是在同一个目录下,所以可以直接使用这个文件中的方法和变量等...
	//但是编译或运行的时候要这样:go build,而不用指出文件名go build main.go
	fmt.Println(subpack.VERSION)
}
