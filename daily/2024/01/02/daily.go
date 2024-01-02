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
		fmt.Println("读取数据成功", data)
	} else {
		fmt.Println("读取数据失败", err)
	}
	return data, err
}

func main() {
	// st := []int{2, 7, 4, 1, 8, 1}
	st, _ := get_task_data(task_file)
	fmt.Println("初始数据", st)

	sort.Sort(sort.IntSlice(st))
	debug_log("第一次排序后的内容", st)

	times := 0
	for len(st) > 1 {
		times = times + 1
		// 判断最大的数字 x 与 第2大的数字 y 是否相同
		// x := st[len(st)-1]
		// y := st[len(st)-2]
		debug_log("第", times, "轮 : x =", st[len(st)-1], ", y =", st[len(st)-2])
		if st[len(st)-1] > st[len(st)-2] {
			// 不相同，则第2大的数字变更为 x - y
			st[len(st)-2] = st[len(st)-1] - st[len(st)-2]
			st = st[:len(st)-1]
			// 重新排序
			sort.Sort(sort.IntSlice(st))
			debug_log("第", times, "轮 x > y")
		} else {
			// 如相同，则 2 个数字全都清除
			st = st[:len(st)-2]
			debug_log("第", times, "轮 x == y")
		}
		debug_log("第", times, "轮操作后的数组", st)
	}

	// fmt.Println("结果", st)
	if len(st) == 0 {
		fmt.Println("最终没有石头剩下")
	} else {
		fmt.Println("最后剩下的石头重量为", st[0])
	}

}
