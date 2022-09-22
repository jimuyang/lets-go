package main

import (
	"fmt"
)

// 数组的大小是固定的，但切片更为灵活

func main19() {
	intArr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(intArr)

	// 一样的 左闭右开
	slice := intArr[1:4]
	fmt.Println(slice)

	// 切片不存储任何数据，只是描述了底层数组的一段
	// 更改切片会反应到数组上
	slice[0] = 9
	fmt.Println(slice)
	fmt.Println(intArr)

	// 切片允许理解成没有长度的数组
	s := []struct {
		i int
		b bool
	}{{2, true}, {2, false}, {1, false}}

	fmt.Println(s)

	// 切片的长度和容量
	// 长度：包含的元素个数
	// 容量：从它的第一个元素开始，到底层数组的结尾的元素个数

	s1 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s1)

	s1 = s1[:0]
	printSlice(s1)

	s1 = s1[:4]
	printSlice(s1)

	s1 = s1[2:]
	printSlice(s1)

	s1 = s1[:4]
	printSlice(s1)

	// 使用make

	s2 := make([]int, 5)
	printSlice(s2) // len=5 cap=5 [0 0 0 0 0]
	s2 = append(s2, 1)
	printSlice(s2) // len=6 cap=10 [0 0 0 0 0 1]

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// fmt.Println("", s)
}
