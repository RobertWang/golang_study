package main

import (
	"sync"
	"time"
)

func demo01() {
	var wg sync.WaitGroup
	foo := make(chan int)
	bar := make(chan int)
	wg.Add(1)
	go func() {
		bar <- 12
	}()
	go func() {
		time.Sleep(time.Duration(1000) * time.Millisecond)
		defer wg.Done()
		select {
		case t := <-bar:
			foo <- t
		default:
			println("default")
		}
	}()
	wg.Wait()
}

func fix_demo01() {
	var wg sync.WaitGroup
	foo := make(chan int, 1) // fixed
	bar := make(chan int, 1) // fixed
	wg.Add(1)
	go func() {
		bar <- 12
	}()
	go func() {
		time.Sleep(time.Duration(1000) * time.Millisecond)
		defer wg.Done()
		select {
		case t := <-bar:
			foo <- t
		default:
			println("default")
		}
	}()
	wg.Wait()
}

func origin() {
	var wg sync.WaitGroup
	foo := make(chan int)
	bar := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case foo <- <-bar:
		default:
			println("default")
		}
	}()
	wg.Wait()
}

func fix_origin() {
	var wg sync.WaitGroup
	foo := make(chan int)
	bar := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case v := <-bar:
			foo <- v
		// case foo <- <-bar:
		default:
			println("default")
		}
		// close(foo)
		// close(bar)
	}()
	wg.Wait()
}

func main() {
	// fix_demo01()
	fix_origin()
}
