package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	/*
		cannot use (Student literal) (value of type Student) as
		People value in variable declaration: Student does not
		implement People (method Speak has pointer receiver)
		compiler(InvalidIfaceAssign)
	*/
	// var peo People = Student{}
	// 改成以下形式
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}
