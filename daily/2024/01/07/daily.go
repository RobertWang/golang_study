package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// 定义数据结构
type CandyData struct {
	Candies      []int `json:"candies"`
	ExtraCandies int   `json:"extraCandies"`
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
func get_task_data(filename string) (data CandyData, err error) {
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

// 实例 1: candies: [4,2,1,1,2], extra: 1
// 实例 2: candies: [12,1,12], extra: 10
// 实例 3: candies: [2,3,5,1,3], extra: 3
// 入口
func main() {
	// 所有示例测试
	// all_candies := []CandyData{
	// 	{Candies: []int{4,2,1,1,2}, ExtraCandies: 1},
	// 	{Candies: []int{12,1,12}, ExtraCandies: 10},
	// 	{Candies: []int{2,3,5,1,3}, ExtraCandies: 3},
	// }
	// for _, data := range all_candies {
	// 	fmt.Println("初始数据", data)
	// 	solution(data.Candies, data.ExtraCandies)
	// }

	// data := CandyData{Candies:[]int{4,2,1,1,2}, ExtraCandies:1}
	data, _ := get_task_data(task_file)
	fmt.Println("初始数据", data)

	// max 函数需要 golang v 1.21 版本
	// fmt.Printf("最大糖果数: %v\n", max(data.Candies))
	solution(data.Candies, data.ExtraCandies)
}

// 解决方案
func solution(candies []int, extraCandies int) {
	// 深拷贝
	bak := make([]int, len(candies))
	copy(bak, candies)
	// 浅拷贝
	// bak := candies
	sort.Ints(bak)
	// fmt.Printf("candy pointer: %p, %v\n", candies, candies)
	// fmt.Printf("bak pointer: %p, %v\n", bak, bak)
	debug_log("最多的糖果数量:", bak[len(bak)-1])
	for inx, candy := range candies {
		fmt.Printf("%d. %d + %d = %d", inx, candy, extraCandies, candy+extraCandies)
		if candy+extraCandies >= bak[len(bak)-1] {
			fmt.Println(" 分配后糖果数量是最多 (true)")
		} else {
			fmt.Println(" 分配后糖果数量不是最多 (false)")
		}
	}
}
