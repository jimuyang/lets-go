package main

import (
	"fmt"
	"io"
	"os"
)

// CopyFile 使用defer来优雅的关闭资源
func CopyFile1(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func testDefer() (result int) {
	// 1. defer的函数调用 参数立即求值
	i := 0
	defer fmt.Println(i) // 0
	i++

	// 栈 后进先出
	for i := 0; i < 4; i++ {
		defer fmt.Print(i) // 3210
	}

	// return前defer还可以做处理
	defer func() { result++ }()
	return 1 // return 2
}

// panic: stops the ordinary flow of control and begins panicking
// recover:

// panic用于爆出一个严重的程序错误 以至于需要程序立刻停止并开始返回
// panic会被层层抛出，但我们可以在defer中recover来接受可能出现的panic并加以处理

// func f() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("recovered in f.", r)
// 		}
// 	}()
// 	fmt.Println("calling g.")
// 	g(0)
// 	// 在我测试时 这一行并不会被打印
// 	fmt.Println("returned from g.")
// }
// func g(i int) {
// 	if i > 3 {
// 		fmt.Println("Panicking")
// 		// 当i为4时 认为程序遇到了严重错误 立即退出
// 		panic(fmt.Sprintf("%v", i))
// 	}
// 	defer fmt.Println("defer in g.", i)
// 	fmt.Println("print in g.", i)
// 	g(i + 1)
// }

// func main10() {
// 	testDefer()
// 	fmt.Println("------")
// 	f()
// }
