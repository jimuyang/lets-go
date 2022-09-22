package main

import (
	"fmt"
	"strconv"
)

type Duck interface {
	Eat()
	Walk()
}

type Dog struct {
}

func (d *Dog) Eat() {
	fmt.Println("dog eat")
	fmt.Println(d == nil)
}

func (d *Dog) Walk() {
	fmt.Println("dog walk")
	fmt.Println(d == nil)
}

func main() {
	i, err := strconv.ParseInt("9473601459061721914", 10, 64)
	fmt.Println(i, err)

	var dog *Dog
	dog.Eat()
	dog.Walk()

	var duck Duck
	// interface像是一种包装类型
	duck = dog               // 这里是将dog放入duck的 动态值  dynamic value中
	fmt.Println(duck == nil) // false duck有了动态值 本身不为nil  体验上允许使用 duck.Walk()了
	duck.Walk()              // 这里执行的时候 Walk的接受者还是dog 为nil
	duck.Eat()
}
