package main

// 闭包
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// 函数式编程的累加器写法
type iAdder func(int) (int, iAdder)

func adder1(base int) iAdder {
	return func(i int) (int, iAdder) {
		// 注意这里的递归
		return base + i, adder1(base + i)
	}
}

func fibnacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// func main() {
// 	// add := adder()
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println(add(i))
// 	// }
// 	iAdder := adder1(0)
// 	for i := 0; i < 10; i++ {
// 		var sum int
// 		sum, iAdder = iAdder(i)
// 		fmt.Println(sum)
// 	}
// 	// fib := fibnacci()
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println(fib())
// 	// }
// }
