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

// 实例 1: [3,4,5,2]
// 实例 2: [1,5,4,5]
// 实例 3: [3,7]
// 入口
func main() {
	// 所有示例测试
	// all_nums := [][]int{
	// 	{3, 4, 5, 2},
	// 	{1, 5, 4, 5},
	// 	{3, 7},
	// }
	// for _, nums := range all_nums {
	// 	fmt.Println("初始数据", nums)
	// 	solution(nums)
	// 	solution_b(nums)
	// }

	// nums := []int{3,4,5,2}
	nums, _ := get_task_data(task_file)
	fmt.Println("初始数据", nums)

	solution(nums)
	// solution_b(nums)
}

// 解决方案
func solution(nums []int) {
	sort.Ints(nums)
	x, y := nums[len(nums)-1], nums[len(nums)-2]
	fmt.Printf("方法1: 最大两个数各自减1后乘积为: (%d - 1) * (%d - 1) = %d * %d = %d\n", x, y, x-1, y-1, (x-1)*(y-1))
}

func solution_b(nums []int) {
	if len(nums) < 2 {
		fmt.Println("数组长度必须大于等于2")
		return
	}
	x, y := nums[0], nums[1]
	if x < y {
		x, y = y, x
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > x {
			x, y = nums[i], x
		} else if nums[i] < y {
			y = nums[i]
		}
	}
	fmt.Printf("方法2: 最大两个数各自减1后乘积为: (%d - 1) * (%d - 1) = %d * %d = %d\n", x, y, x-1, y-1, (x-1)*(y-1))
}
