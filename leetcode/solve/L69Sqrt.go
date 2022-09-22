package main

import (
	"math"
)

func mySqrt(x int) int {
	var f float64
	f = 1.0
	var r, r1 int
	r = 1
	for {
		f = f - (f*f-float64(x))/(2*f)
		r1 = int(math.Floor(f))
		if r1 != r {
			r = r1
		} else {
			return r
		}
	}
}

// func main() {
// 	fmt.Println(mySqrt(8))
// }
