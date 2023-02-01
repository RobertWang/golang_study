package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {

	list = make(map[string]Student)

	student := Student{"Aceld"}

	list["student"] = student
	// 下面这行会报错
	// list["student"].Name = "LDB"

	fmt.Println("origin source: ", list["student"])

	// solution 1:
	var list1 map[string]Student
	list1 = make(map[string]Student)
	student1 := Student{"Aceld"}
	list1["student"] = student1
	tmpStudent := list1["student"]
	tmpStudent.Name = "LDB"
	list["student"] = tmpStudent
	fmt.Println("solution 1: ", list["student"])

	// solution 2:
	var list2 map[string]*Student
	list2 = make(map[string]*Student)
	student2 := Student{"Aceld"}
	list2["student"] = &student2
	list2["student"].Name = "LDB"
	fmt.Println("solution 2: ", list["student"])
}
