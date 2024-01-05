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
func get_task_data(filename string) (data [][]int, err error) {
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

// 实例 1: [[1,2]]
// 实例 2: [[1,3],[2,3]]
// 实例 3: [[1,3],[2,3],[3,1]]
// 实例 4: [[1,3],[1,4],[2,3],[2,4],[4,3]]
// 入口
func main() {
	// persons := [][]int{{1,3},{1,4},{2,3},{2,4},{4,3}}
	persons, _ := get_task_data(task_file)
	fmt.Println("初始数据", persons)

	solution(persons)
}

// 解决方案
func solution(persons [][]int) {
	judge := -1
	var in, out, all map[int]int
	in = make(map[int]int)
	out = make(map[int]int)
	all = make(map[int]int)
	for _, rel := range persons {
		if _, ok := all[rel[0]]; !ok {
			all[rel[0]] = 0
		}
		if _, ok := all[rel[1]]; !ok {
			all[rel[1]] = 0
		}
		if _, ok := in[rel[0]]; ok {
			out[rel[0]]++
		} else {
			out[rel[0]] = 0
		}
		if _, ok := in[rel[1]]; ok {
			in[rel[1]]++
		} else {
			in[rel[1]] = 1
		}
	}
	person_count := len(all)
	for person, _ := range all {
		debug_log("当前是", person)
		in_count, out_count := 0, 0
		out_count, _ = out[person]
		in_count, _ = in[person]
		if out_count == 0 && in_count == person_count-1 {
			// fmt.Println("法官是", person)
			judge = person
			break
		}
	}
	if judge == -1 {
		fmt.Println("结果: 小镇没有法官")
	} else {
		fmt.Println("结果: 法官是", judge)
	}
}
