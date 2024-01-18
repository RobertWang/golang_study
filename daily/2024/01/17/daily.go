package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"strconv"
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

// 实例 1: 2024-03-01
// 入口
func main() {
	// data := "2024-03-01"
	data, _ := get_task_data(task_file)

	fmt.Println("初始数据", data)

	solution(data)
}

// 解决方案
func solution(date string) {
	mdays := []int{
		31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31,
	}
	sp := strings.Split(date, "-")
	y, ok := strconv.Atoi(sp[0])
	if ok != nil {
		panic("month error")
	}
	m, ok := strconv.Atoi(sp[1])
	if ok != nil {
		panic("month error")
	}
	d, ok := strconv.Atoi(sp[2])
	if ok != nil {
		panic("day error")
	}

	if y % 400 == 0 || (y % 4 == 0 && y % 100 != 0) {
		debug_log(y, "是闰年")
		mdays[1] = 29
	}
	total := d
	debug_log(sp, y, m, d)
	debug_log(mdays)
	for i, t := range mdays[:m-1] {
		debug_log(i, t)
		total += t
	}
	fmt.Println(date, "是", sp[0], "年的第", total, "天")
}
