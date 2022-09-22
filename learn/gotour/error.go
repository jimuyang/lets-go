package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(), "npe",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// 这样写会导致死循环 看调用链路：Sprint->doPrint->printArg->handleMethods->Error()
	// 简而言之：在print一个值时，如果是error 会调用Error()的
	// return fmt.Sprintf("cannot Sqrt negative number: %v", e)
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt1(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func main10() {
	fmt.Println(Sqrt1(2))
	fmt.Println(Sqrt1(-2))
}

// func main() {
// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 	}
// }
