package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort" // 用于 方案 2 : 暴力解决方案
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

// 入口
func main() {
	// heights := []int{2, 7, 4, 1, 8, 1}
	heights, _ := get_task_data(task_file)
	fmt.Println("初始数据", heights)

	solution(heights)
	// solution_b(heights, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 解决方案
func solution(heights []int) {
	left, right := 0, len(heights)-1
	max_vol := 0
	for left < right {
		vol := min(heights[left], heights[right]) * (right - left)
		if vol > max_vol {
			max_vol = vol
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	fmt.Println("最大容量", max_vol)
}

// 方案 2 : 暴力解决
type Volume struct {
	Value int
	Left  int
	Right int
}

func solution_b(heights []int, inx_args ...int) {
	inx := -1
	if len(inx_args) > 0 {
		inx = inx_args[0]
	}

	left, right := 0, len(heights)-1
	vols := make([]Volume, 0)
	done := false

	for done == false {
		debug_log("left", left, ", right", right, ", done", done)
		if left >= right {
			done = true
			break
		}

		x := heights[left]
		y := heights[right]
		vol := min(x, y) * (right - left)
		current := &Volume{
			Left:  left,
			Right: right,
			Value: vol,
		}
		vols = append(vols, *current)
		debug_log("本轮", "x", x, "y", y, "result:", current)
		if x < y {
			debug_log("left 指针右移")
			left++
		} else {
			debug_log("right 指针左移")
			right--
		}
	}

	// 进行一次排序
	sort.Slice(vols, func(i, j int) bool {
		return vols[i].Value > vols[j].Value
	})

	// 显示结果
	if inx == -1 {
		for i, v := range vols {
			fmt.Println("第", i+1, "个结果", "容量:", v.Value)
			fmt.Println("    位置: (", v.Left, ", ", v.Right, ") 长度:", v.Right-v.Left)
			fmt.Println("    高度: (", heights[v.Left], ", ", heights[v.Right], ") 最短高度:", min(heights[v.Left], heights[v.Right]))
		}
	} else {
		v := vols[inx]
		fmt.Println("容量最大的结果", v.Value, "位置: (", v.Left, ", ", v.Right, ") 长度:", v.Right-v.Left, "高度: (", heights[v.Left], ", ", heights[v.Right], ") 最短高度:", min(heights[v.Left], heights[v.Right]))
	}
}
