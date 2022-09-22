package main

import (
	"fmt"
)

func main0() {
	// go中的数组不能改变大小
	var a [2]string
	fmt.Println(a)

	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a[0], a[1])

	intArr := [...]int{1, 2, 2, 4546}
	fmt.Println(intArr)
}
