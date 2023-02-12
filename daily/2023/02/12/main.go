package main

import "reflect"

// 只有接口才允许类型断言
// 通过反射获取类型
func GetValue() int {
	return 1
}

// 改成如下形式，函数返回接口类型，再进行类型断言
func GetValueFix() interface{} {
	return 1
}

// 基于原始函数进行类型判断
func origin_fix1() {
	i := GetValue()
	itype := reflect.TypeOf(i).String()
	println("origin fix 1: ", itype)
	switch itype {
	case "int":
		println("int")
	case "string":
		println("string")
	case "interface":
		println("interface")
	default:
		println("unknown")
	}

}

// 通过原始函数修正进行类型判断
func origin_fix2() {
	i := GetValueFix()
	println("origin fix 2: ")
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

}

func main() {
	// 修正方案1
	origin_fix1()
	// 修正方案2
	origin_fix2()
}
