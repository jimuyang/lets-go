package main

import (
	"fmt"
	"math"
)

// 函数也是值。它们可以像其它值一样传递。

// 函数值可以用作函数的参数或返回值。

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main11() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	adder1, adder2 := adder(), adder()
	for i := 0; i < 10; i++ {
		adder1(i)
		adder2(2 * i)
	}
	fmt.Println(adder1(0))
	fmt.Println(adder2(0))

	fn := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}

// 函数可以是一个闭包
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	i := 1
	sum := 0
	return func() int {
		t := sum
		sum = sum + i
		i = t
		return sum
	}
}
