package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	// "strconv"
	// "strings"
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

type PasswordData struct {
	Password string `json:"password"`
	Target   int    `json:"target"`
}

// 从文件中获取任务数据
func get_task_data(filename string) (data PasswordData, err error) {
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

// 实例 1: PasswordData{Password:"s3cur1tyC0d3",Target:4}
// 入口
func main() {
	// data := PasswordData{Password:"s3cur1tyC0d3",Target:4}
	data, _ := get_task_data(task_file)
	fmt.Println("初始数据", data)

	solution(data.Password, data.Target)
}

// 解决方案
func solution(password string, target int) {
	result := password[target:] + password[:target]
	fmt.Println(result)
}
