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
func get_task_data(filename string) (data [][]string, err error) {
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

// 入口
func main() {
	// 1: [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
	// 2: [["B","C"],["D","B"],["C","A"]]
	// 3: [["A","Z"]]
	// data := [][]string{
	// 	{"London", "New York"},
	// 	{"New York", "Lima"},
	// 	{"Lima", "Sao Paulo"},
	// }

	data, _ := get_task_data(task_file)
	fmt.Println("初始数据", data)

	solution(data)
}

// 解决方案
func solution(data [][]string) {
	starts := make(map[string]int)
	for inx, points := range data {
		starts[points[0]] = inx
	}
	debug_log("所有起始城市", starts)
	exists := true
	for _, points := range data {
		_, exists = starts[points[1]]
		if !exists {
			fmt.Println("终归点城市是", points[1])
		}
	}
	if exists {
		fmt.Println("当前城市路径规划没有终归")
	}
}
