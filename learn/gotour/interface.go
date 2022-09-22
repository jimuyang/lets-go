package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 接口类型是由一组方法集合定义的类型
// 接口类型的变量可以保存任何实现了这些方法的值

// type Abser interface {
// 	Abs() float64
// }

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f > 0 {
// 		return float64(f)
// 	}
// 	return float64(-f)
// }

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main1() {
// 	var a Abser

// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}

// 	a = f
// 	a = &v
// 	// a = v

// 	fmt.Println(a.Abs())
// }

// 接口也是值 它们可以像其他值一样传递

// 在内部，接口值可以看做包含值和具体类型的元组：

// (value, type)
// 接口值保存了一个具体底层类型的具体值。

// 接口值调用方法时会执行其底层类型的同名方法。

// 类型断言 : 提供了访问接口底层具体值的方式
// t := i.(T) // 这里断言i这个接口底层保存的是实际类型为T的值 并将实际值赋值给t 如果底层不是T类型 就是触发panic
// 不会触发panic的方式 t, ok := i.(T)  当底层不是T类型时 ok为false t为T类型零值

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main2() {
	do(21)
	do("hello")
	do(true)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main3() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

// 练习 打印ip
type IPAddr [4]byte

func (ip IPAddr) String() string {
	// var b strings.Builder
	// b.WriteByte('a')
	// b.WriteByte('a')
	// b.WriteByte('a')
	// fmt.Println(b.String())

	// str := ""
	// for _, e := range [4]byte(ip) {
	var strs [4]string
	for i := 0; i < 4; i++ {
		strs[i] = strconv.Itoa(int(ip[i]))
	}
	return strings.Join(strs[:], ".")
}
func main4() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type MyType struct {
	name string
}

func (myType *MyType) String() string {
	return myType.name
}

func main14() {
	fmt.Println(&MyType{"yangfan"})
}
