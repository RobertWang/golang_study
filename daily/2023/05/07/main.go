package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

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

func main() {
	var m map[string]int
	var n map[string]int

	fmt.Println(m == nil)
	fmt.Println(n == nil)

	// 不能通过编译
	//fmt.Println(m == n)

	fmt.Println(reflect.DeepEqual(m, n))

	cm1 := CompareMap((map[string]interface{})m, n)
	fmt.Println(cm1)

}
