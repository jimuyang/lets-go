package main

import (
	"fmt"
	"sync"
	"time"
)

func say1(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// 信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
func sum(s []int, c chan int) {
	sum := 0
	for _, i := range s {
		sum += i
	}
	c <- sum // 将sum送入chan
}

func main8() {
	// go say("world")
	// say("hello")

	c := make(chan int)
	s := []int{7, 2, 8, -9, 4, 0}
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从c中接收
	fmt.Println(x, y)

	// chan可以带缓冲
	// 仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch) // all goroutines are asleep - deadlock!

	// 循环 for i := range c 会不断从信道接收值，直到它被关闭。

	c1 := make(chan int, 10)
	go fibonacci1(cap(c1), c1)
	for i := range c1 {
		fmt.Println(i)
	}
}

func fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("put", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("receive quit")
			return
		}
	}
}

func main5() {
	c := make(chan int, 1)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

// go中的锁 sync.Mutex

// SafeCounter ..
type SafeCounter struct {
	count int
	mux   sync.Mutex
}

func (c *SafeCounter) incAndGet() int {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.count++
	return c.count
}

func (c *SafeCounter) unsafe() int {
	// c.mux.Lock()
	// defer c.mux.Unlock()
	c.count++
	return c.count
}

func main9() {
	c := SafeCounter{count: 0}
	for i := 0; i < 1000; i++ {
		// go c.incAndGet()
		go c.unsafe()
	}
	time.Sleep(time.Second)
	fmt.Println(c.count)
}
