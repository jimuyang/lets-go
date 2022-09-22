package main

import (
	"fmt"
	"sync"
	"time"
)

// go run -race goroutine.go
// goroutine切换的时间点
// I/O select channel 锁 函数调用（有时） runtime.Gosched()
func main111() {
	// var a [10]int
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		for {
	// 			// fmt.Println("hello from i: ", i)
	// 			a[i]++
	// 			runtime.Gosched()
	// 		}
	// 	}(i)
	// }
	// time.Sleep(time.Millisecond)
	// fmt.Println(a)
	// fmt.Println("main finished")

	// for i := 0; i < 1000; i++ {
	// 	go func(i int) {
	// 		for {
	// 			fmt.Println("hello from i: ", i)
	// 		}
	// 	}(i)
	// }
	// time.Sleep(time.Minute)
	// fmt.Println("main finished")
	// beforePrintA()

	twoPrint()
	time.Sleep(time.Millisecond)
}

var a string

func printA() {
	fmt.Println(a)
}

func beforePrintA() {
	a = "hello"
	go printA()
}

var i int
var done bool
var once sync.Once

func initI() {
	i = 2
	done = true
}

func doPrint() {
	// 这个dubble-check因可能的重排序是错误的
	if !done {
		once.Do(initI)
	}
	fmt.Println(i)
}

func twoPrint() {
	for i := 0; i < 10; i++ {
		go doPrint()
	}
}
