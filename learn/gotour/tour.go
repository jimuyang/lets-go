package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

// 函数外的每个语句都必须以关键字开始 如 var func等等

// func main() {

// 	var i, j int = 1, 2
// 	k := 3

// 	c, python, java := true, false, "!no"
// 	fmt.Println(i, j, k, c, python, java)

// 	zero_type()
// 	type_conversion()
// 	type_inference()
// 	test_for()
// 	test_if()
// 	// Sqrt(2)
// 	test_switch()
// 	test_defer()
// 	test_pointer()
// 	testStruct()

// 	test_defer_stack()

// 	newton(8)
// }

// go的基本类型
// bool string int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr
// byte :uint8
// rune :int32 表示一个Unicode码点
// float32 float64
// complex64 complex128

func zeroType() {
	var i int
	var f float32
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// 类型推导
func type_inference() {
	v := 42
	fmt.Printf("v is of type: %T\n", v)
}

// go只有for循环
func test_for() {
	var sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum <= 100 {
		sum++
	}
	fmt.Println(sum)
}

func test_if() {
	var x = 10
	if x < 10 {
		fmt.Println("x < 10")
	} else if x > 10 {
		fmt.Println("x > 10")
	} else {
		fmt.Println("x = 10")
	}
}

// switch是编写一连串 if-else的简便写法
// 重要不同：go中的break是默认提供的，反而除非以 fallthrough结束

func test_switch() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// defer 新东西
func test_defer() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}

func test_defer_stack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done!")
}

// pointer go拥有指针
// 指针保存了值的内存地址。
func test_pointer() {
	fmt.Println("pointer use:")
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

// // struct
// type Vertex struct {
// 	X, Y int
// }

// func testStruct() {
// 	var (
// 		v1 = Vertex{1, 2}
// 		v2 = Vertex{X: 1}
// 		v3 = Vertex{}
// 		p  = &Vertex{Y: 2}
// 	)
// 	fmt.Println(v1, v2, v3, p)
// }

func newton(x float64) float64 {
	fmt.Println("newton start")
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

// go switch很特别 默认提供了break 而需要写fallthrough

// defer语句会将函数推迟到外层函数返回之后执行
// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用
// 推迟的函数调用会被压入一个栈中 当外层函数返回时，被推迟的函数会按照 后进先出 顺序调用

// CopyFile 使用defer来安全的释放资源
func CopyFile(dstName, srcName string) (written int64, err error) {
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

func deferUseReturn() (i int) {
	defer func() { fmt.Println(i) }()
	return 1
}

// func main() {
// 	deferUseReturn()
// }

//This is convenient for modifying the error return value of a function; we will see an example of this shortly.

// Panic Recover

// func main() {
// 	f()
// 	fmt.Println("return normally from f.")
// }

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}
	}()
	fmt.Println("calling g")
	g(0)
	fmt.Println("return normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("defer in g", i)
	fmt.Println("print in g", i)
	g(i + 1)
}

func zeroValue() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q end", i, f, b, s)
	fmt.Printf("%v %v %v %v end", i, f, b, s)
}

// func main() {
// 	zeroValue()

// 	// 常量申明 不可以使用 := 语法
// 	const World = "世界"
// }

// for循环的range 可以遍历切片或者映射
// for 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。

func testRange() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for _, v := range pow {
		fmt.Println(v)
		v = 1
	}
	fmt.Println(pow)
}

func main() {
	// testRange()

}
