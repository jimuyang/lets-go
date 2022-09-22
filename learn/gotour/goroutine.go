package main

import (
	"fmt"
	"time"
)

// Go 程（goroutine）是由 Go 运行时管理的轻量级线程。
// go f(x, y, z) // 会启动一个新的 Go 程并执行 f
// f, x, y 和 z 的求值发生在当前的 Go 程中，而 f 的执行发生在新的 Go 程中
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main12() {
	go say("world")
	say("hello")
}
