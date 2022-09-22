package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Communication Sequential Process CSP

// don't communicate by sharing memory, share memory by communicating
// 不要通过共享内存来通信，通过通信来共享内存

func worker(id int, c chan int) {
	for {
		// n, ok := <-c
		// if !ok {
		// 	break
		// }
		for n := range c {
			fmt.Printf("worker: %d receive: %c\n", id, n)
		}
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func channelDemo() {
	var c chan int // c == nil
	fmt.Println(c)

	// 无缓冲的channel
	c = make(chan int)
	for i := 0; i < 3; i++ {
		go worker(i, c)
	}
	for i := 0; i < 100; i++ {
		c <- i
	}
	time.Sleep(time.Millisecond)

	// var channels [10]chan<- int
	// for i := 0; i < 10; i++ {
	// 	channels[i] = createWorker(i)
	// }
	// for i := 0; i < 10; i++ {
	// }

	// 有缓冲的channel
	// c = make(chan int, 5)
	// c <- 1
	// c <- 2
	// n := <-c
	// fmt.Println(n)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	// close(c)
	time.Sleep(time.Millisecond)
}

type doneWorker struct {
	work chan int
	done chan bool
}

func createDoneWorker(id int) doneWorker {
	worker := doneWorker{
		work: make(chan int),
		done: make(chan bool),
	}
	go func(id int, work chan int, done chan bool) {
		for w := range work {
			fmt.Printf("worker %d is doing %d\n", id, w)
			// go func() { done <- true }() // 但这样会让协程数量翻倍
			done <- true
		}
	}(id, worker.work, worker.done)
	return worker
}

func useDoneWorker() {
	var workers [10]doneWorker
	for i := 0; i < 10; i++ {
		workers[i] = createDoneWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].work <- i
		// <-workers[i].done 这样是顺序执行
	}
	// wait all done
	for i := 0; i < 10; i++ {
		<-workers[i].done
	}
	for i := 0; i < 10; i++ {
		workers[i].work <- i + 10
	}
	// wait all done
	for i := 0; i < 10; i++ {
		<-workers[i].done
	}
}

func useWaitGroup() {
	var workers [10]doneWorker
	for i := 0; i < 10; i++ {
		workers[i] = createDoneWorker(i)
	}

	var wg sync.WaitGroup
	wg.Add(20)
	go func() {
		for {
			for i := 0; i < 10; i++ {
				<-workers[i].done
				wg.Done()
			}
		}
	}()

	for i := 0; i < 10; i++ {
		workers[i].work <- i
	}
	for i := 0; i < 10; i++ {
		workers[i].work <- i + 10
	}
	wg.Wait()
}

func channelCommunicate() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

// select
func useSelect() {
	// var c1, c2 chan int // c1 c2 = nil, nil 对于nil channel select是阻塞的
	c1, c2 := genChannelAndSend(), genChannelAndSend()
	finish := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		select {
		case n := <-c1:
			fmt.Println("receive from c1:", n)
		case n := <-c2:
			fmt.Println("receive from c2:", n)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("1 second") // 这个定时会影响到上面的timeout
		case <-finish:
			fmt.Println("finish")
			return
			// default: // 使用default会让select变为非阻塞
		}
	}
}

func genChannelAndSend() chan int {
	c := make(chan int)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			c <- i
		}
	}()
	return c
}

func main2() {
	// channelDemo()
	// bufferedChannel()
	// channelCommunicate()
	// useDoneWorker()
	// useWaitGroup()
	useSelect()
}
