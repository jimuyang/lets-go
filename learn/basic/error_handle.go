package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	for i := 3; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println(10)
}

// func main() {
// 	tryDefer()
// }

func tryOpenFile() {
	file, err := os.Open("abc.txt")
	defer file.Close()

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic("unknown error" + err.Error())
		} else {
			fmt.Println(pathError.Err)
		}
	}
	all, err := ioutil.ReadAll(file)
	fmt.Println(all)
}
