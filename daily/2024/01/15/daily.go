package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// 任务数据文件
const task_file string = "task.json"

// 是否输出调试信息
const debug_mode bool = false

// 输出调试信息
func debug_log(info ...interface{}) {
	if debug_mode == true {
		fmt.Println("[DEBUG]", info)
	}
}

// 从文件中获取任务数据
func get_task_data(filename string) (data string, err error) {
	file, _ := os.Open(filename)
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&data); err == nil {
		debug_log("读取数据成功", data)
	} else {
		debug_log("读取数据失败", err)
	}
	return data, err
}

// 实例 1: "a.aef.qerf.bb"
// 入口
func main() {
	// data := "a.aef.qerf.bb"
	data, _ := get_task_data(task_file)
	fmt.Println("初始数据", data)

	solution(data)
}

// 解决方案
func solution(password string) {
	/*
		### 语法

		`func Replace(s, old, new string, n int) string`

		### 参数

		参数 | 描述
		--  | --
		s   | 要替换的整个字符串。
		old | 要替换的字符串。
		new | 替换成什么字符串。
		n   | 要替换的次数，-1，那么就会将字符串 s 中的所有的 old 替换成 new。

	*/

	encoded := strings.Replace(password, ".", " ", -1)
	fmt.Println("加密后的字符串为:", encoded)
}
