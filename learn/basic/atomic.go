package main

import (
	"fmt"
	"sync"
	"time"
)

// 传统的同步机制
// waitGroup Cond Mutex

type atomicInt struct {
	val  int
	lock sync.Mutex
}

func (ai *atomicInt) increment() {
	fmt.Println("safe")
	func() {
		ai.lock.Lock()
		defer ai.lock.Unlock()
		ai.val++
	}()
}

func (ai *atomicInt) get() int {
	ai.lock.Lock()
	defer ai.lock.Unlock()

	return ai.val
}

// go run -race learn/basic/atomic.go

func main5() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
