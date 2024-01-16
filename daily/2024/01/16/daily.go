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
func get_task_data(filename string) (data int, err error) {
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

// 实例 1: 5
// 入口
func main() {
	// data := 5
	data, _ := get_task_data(task_file)
	fmt.Println("初始数据", data)

	solution(data)
}

// 解决方案
func solution(num int) {
	var series []interface{}
	var i int
	var item interface{}

	for i = 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			item = "FizzBuzz"
		} else if i%3 == 0 {
			item = "Fizz"
		} else if i%5 == 0 {
			item = "Buzz"
		} else {
			item = i
		}
		series = append(series, item)
	}
	fmt.Println("结果序列为:", series)
}
