package main

import (
	"fmt"
	"sync"
)

var (
	x    int
	lock sync.Mutex
)

func Print(i int) {
	fmt.Println("hello routine", i)
}
func routine() {
	var wg sync.WaitGroup // 声明一个waitgroup
	wg.Add(5)             // 添加五个任务
	for i := 1; i <= 5; i++ {
		go func(j int) {
			defer wg.Done() // 完成任务任务数-1
			Print(j)
		}(i)
	}
	wg.Wait() // 等待所有任务完后
}
func CalSquare() {
	src := make(chan int)
	dest := make(chan int, 3)
	go func() {
		defer close(src)
		for i := 1; i <= 10; i++ {
			src <- i
		}
	}()
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()
	for i := range dest {
		fmt.Println(i)
	}
}
func addwithlock() {
	for i := 1; i <= 2000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
}
func addwithoutlock() {
	for i := 1; i <= 2000; i++ {
		x++
	}
}
func main() {
	routine()
}
