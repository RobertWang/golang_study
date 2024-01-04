package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
func get_task_data(filename string) (data []int, err error) {
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

func main() {
	// stones := []int{2, 7, 4, 1, 8, 1}
	stones, _ := get_task_data(task_file)
	fmt.Println("初始数据", stones)

	solution(stones)
}

// 解决方案
func solution(stones []int) {
	sort.Sort(sort.IntSlice(stones))
	debug_log("第一次排序后的内容", stones)

	times := 0
	for len(stones) > 1 {
		times = times + 1
		// 判断最大的数字 x 与 第2大的数字 y 是否相同
		// x := stones[len(stones)-1]
		// y := stones[len(stones)-2]
		debug_log("第", times, "轮 : x =", stones[len(stones)-1], ", y =", stones[len(stones)-2])
		if stones[len(stones)-1] > stones[len(stones)-2] {
			// 不相同，则第2大的数字变更为 x - y
			stones[len(stones)-2] = stones[len(stones)-1] - stones[len(stones)-2]
			stones = stones[:len(stones)-1]
			// 重新排序
			sort.Sort(sort.IntSlice(stones))
			debug_log("第", times, "轮 x > y")
		} else {
			// 如相同，则 2 个数字全都清除
			stones = stones[:len(stones)-2]
			debug_log("第", times, "轮 x == y")
		}
		debug_log("第", times, "轮操作后的数组", stones)
	}

	// fmt.Println("结果", stones)
	if len(stones) == 0 {
		fmt.Println("最终没有石头剩下")
	} else {
		fmt.Println("最后剩下的石头重量为", stones[0])
	}
}
