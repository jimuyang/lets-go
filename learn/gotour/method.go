package main

import (
	"fmt"
	"math"
)

// go中没有类 但可以定义特殊的 带接受者的参数
type Vertex struct {
	X, Y int
	In   Inner
}
type Inner struct {
	H, Z int
}

func (v Vertex) Abs() int {
	x := float64(v.X)
	y := float64(v.Y)
	return int(math.Sqrt(x*x + y*y))
}

func (v *Vertex) Scale(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func (v Vertex) Test() {
	v.In.Z = 100
	fmt.Printf("in Test: &v = %p \n", &v)
	fmt.Printf("in Test: &v.In = %p \n", &(v.In))
	fmt.Printf("in Test: &v.In.Z = %p \n", &(v.In.Z))
}

func main16() {
	v := Vertex{3, 4, Inner{0, 1}}
	fmt.Println(v)
	fmt.Printf("in main: &v = %p \n", &v)
	fmt.Printf("in main: &v.In = %p \n", &(v.In))
	fmt.Printf("in main: &v.In.Z = %p \n", &(v.In.Z))
	v.Scale(10)
	v.Test()
	fmt.Println(v)
}

// 也可以为非结构体申明方法 不能为内类型声明方法
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// func main() {
// 	f := MyFloat(-math.Sqrt2)
// 	fmt.Println(f.Abs())
// }

// 所以go中如果不写指针 都是值copy传递 不管是不是结构体
