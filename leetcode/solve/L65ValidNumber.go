package main

import (
	"fmt"
	"reflect"
)

func testLoopString() {
	str := "hi, 世界"

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
	fmt.Println(reflect.TypeOf(str[0])) // uint8

	fmt.Println("------")
	for _, ch := range str {
		fmt.Println(ch)
		// fmt.Println(string(ch))
		fmt.Println(reflect.TypeOf(ch)) // int32
	}

	// fmt.Println(rune(8))

}

func isNumber(s string) bool {
	finals := []int{0, 0, 0, 1, 0, 1, 1, 0, 1}
	transferTable := [][]int{
		{0, 1, 6, 2, -1},
		{-1, -1, 6, 2, -1},
		{-1, -1, 3, -1, -1},
		{8, -1, 3, -1, 4},
		{-1, 7, 5, -1, -1},
		{8, -1, 5, -1, -1},
		{8, -1, 6, 3, 4},
		{-1, -1, 5, -1, -1},
		{8, -1, -1, -1, -1},
	}

	state := 0
	for _, ch := range s {
		i := tranferIndex(ch)
		if i < 0 {
			return false
		}
		state = transferTable[state][i]
		if state < 0 {
			return false
		}
	}
	return finals[state] > 0
}

func tranferIndex(ch int32) int {
	switch ch {
	case ' ':
		return 0
	case '+':
		fallthrough
	case '-':
		return 1
	case '.':
		return 3
	case 'e':
		return 4
	default:
		if ch >= '0' && ch <= '9' {
			return 2
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(isNumber("0"))
// }
