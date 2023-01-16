package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x    int
	lock sync.Mutex
)

func Print(i int) {
	fmt.Println("hello routine", i)
}
func routine() {
	for i := 1; i <= 5; i++ {
		go func(j int) {
			Print(j)
		}(i)
	}
	time.Sleep(time.Second)
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
	for i := 1; i <= 5; i++ {
		go addwithlock()
	}
	time.Sleep(time.Second)
	fmt.Println(x)
	x = 0
	for i := 1; i <= 5; i++ {
		go addwithoutlock()
	}
	time.Sleep(time.Second)
	fmt.Println(x)
}
