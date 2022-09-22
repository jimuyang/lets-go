package main

import (
	"fmt"
	"math"
)

/**
 * 使用牛顿法来计算一个平方根函数
 */
func Sqrt(x float64) float64 {
	// 猜测从1开始
	var z = 1.0

	for i := 0; i < 10; i++ {
		fmt.Println(z)
		dz := (z*z - x) / (2 * z)
		if math.Abs(dz) < 0.00001 {
			return z
		}
		z -= dz
	}
	return z
}

// func main() {
// 	Sqrt(2)
// }
