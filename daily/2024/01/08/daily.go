package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// 定义数据结构
type Data struct {
	TimeSeries []int // `json:"timeSeries"`
	Duration   int   // `json:"duration"`
}

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
func get_task_data(filename string) (data Data, err error) {
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

// 实例 1: timeSeries: [1,4], duration: 2
// 实例 2: timeSeries: [1,2], duration: 2
// 实例 3: timeSeries: [1,2,4], duration: 2
// 入口
func main() {
	// 所有示例测试
	// all_data := []Data{
	// 	{TimeSeries: []int{1, 4}, Duration: 2},
	// 	{TimeSeries: []int{1, 2}, Duration: 2},
	// 	{TimeSeries: []int{1, 2, 4}, Duration: 2},
	// }
	// for _, data := range all_data {
	// 	fmt.Println("初始数据", data)
	// 	solution(data.TimeSeries, data.Duration)
	// }

	// data := Data{TimeSeries: []int{1, 2, 4}, Duration: 2}
	data, _ := get_task_data(task_file)
	fmt.Println("初始数据", data)

	solution(data.TimeSeries, data.Duration)
}

// 解决方案
func solution(timeSeries []int, duration int) {
	total, expired := 0, 0
	for _, point := range timeSeries {
		if point > expired {
			total += duration
		} else {
			total += point + duration - expired
		}
		expired = point + duration
	}
	fmt.Println("提莫攻击所致中毒的总时间为", total)
}
