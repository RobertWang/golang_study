package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
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

// 实例 1: "LVIIIMCMXCIV" // 2050
// 入口
func main() {
	// data := "LVIIIMCMXCIV"
	data, _ := get_task_data(task_file)
	fmt.Println("初始数据", data)

	solution(data)
}

// 解决方案
func solution(roma_num string) {
	result := 0
	length := len(roma_num)
	symbol := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}
	for i, char := range roma_num {
		if i < length-1 && symbol[string(char)] < symbol[string(roma_num[i+1])] {
			result -= symbol[string(char)]
		} else {
			result += symbol[string(char)]
		}	
	}
	fmt.Println(roma_num, "转换为整数为", result)
}
