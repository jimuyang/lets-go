package main

import "fmt"

// 分式化简
func fraction(cont []int) []int {
	fraction := []int{cont[len(cont)-1], 1}
	for i := len(cont) - 2; i >= 0; i-- {
		fraction = fractionAdd(cont[i], fraction)
		fmt.Println(fraction[0], "/", fraction[1])
	}
	return fraction
}

func fractionAdd(a int, fraction []int) []int {
	// 倒分数
	fraction[0], fraction[1] = fraction[1], fraction[0]
	fraction[0] += fraction[1] * a
	return fraction
}
