package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 判断两个 map 是否相现
// 这个方法是判断给定的 dm1 是否是 dm2 的真子集的逻辑
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

// EqualMap 比较 map 是否相等
// * 只适用成员为标量的 map
// 这个方法是判断给定的 m1 是否是 m2 的真子集的逻辑
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

// 类型转换
func convertType(in map[string]int) (out map[string]interface{}) {
	out = make(map[string]interface{})
	for k, v := range in {
		out[k] = v
	}
	return
}

func main() {
	var m map[string]int
	var n map[string]int
	fmt.Println("m is nil? ", m == nil)
	fmt.Println("n is nil? ", n == nil)

	m = make(map[string]int)
	n = make(map[string]int)
	m["a"] = 0
	n["a"] = 0
	n["b"] = 1

	fmt.Println("m is nil? ", m == nil)
	fmt.Println("n is nil? ", n == nil)

	// 不能通过编译
	//fmt.Println(m == n)

	c_m1 := convertType(m)
	c_n1 := convertType(n)
	fmt.Println("m[string]interface{} := ", c_m1)
	fmt.Println("n[string]interface{} := ", c_n1)

	// 使用反射中的方法判断
	fmt.Println("内置的反射方法判断", reflect.DeepEqual(m, n))

	cm1 := CompareMap(c_m1, c_n1)
	fmt.Println("CompareMap 方法判断(1)", cm1)
	cm1 = CompareMap(c_n1, c_m1)
	fmt.Println("CompareMap 方法判断(2)", cm1)

	cm2 := EqualMap(c_m1, c_n1)
	fmt.Println("EqualMap 方法判断(1)", cm2)
	cm2 = EqualMap(c_n1, c_m1)
	fmt.Println("EqualMap 方法判断(2)", cm2)

	cm3 := EqualMapReflect(c_m1, c_n1)
	fmt.Println("EqualMapReflect 方法判断", cm3)

}
