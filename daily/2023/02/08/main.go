package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func origin() {
	t := struct {
		time.Time
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	m, _ := json.Marshal(t)
	fmt.Printf("%s", m)
}

func demo1() {
	t := struct {
		Time time.Time
		N    int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}
	m, _ := json.Marshal(t)
	fmt.Printf("%s", m)
}

func demo2() {
	type Tmp struct {
		Time time.Time
		N    int
	}
	t := Tmp{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}
	m, _ := json.Marshal(t)
	fmt.Printf("%s", m)
}

func runTest(fname string, fn func()) (err error) {
	fmt.Printf("%s :\n", fname)
	fn()
	fmt.Printf("\n%s end\n", fname)
	return nil
}

func main() {
	runTest("origin", origin)
	runTest("demo1", demo1)
	runTest("demo2", demo2)
}
