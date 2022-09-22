package main

import (
	"fmt"
)

// go中的指针 指针保存了值的内存地址
// 类型 *T 是指向 T 类型值的指针。其零值为 nil。
func testPtr() {
	i := 42

	p := &i
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)
}

// Node 测试指针
type Node struct {
	value int
	left  *Node
	right *Node
}

func main17() {
	// testPtr()
	node1 := &Node{1, nil, nil}
	node1.left = &Node{2, nil, nil}
	node1.right = &Node{3, nil, nil}

	left := node1.left
	right := node1.right
	fmt.Println(left.value)
	fmt.Println(right.value)
	swap(node1)
	fmt.Println(left.value)
	fmt.Println(right.value)

	change(node1)
	fmt.Println(node1.value)
}

func swap(node *Node) {
	temp := node.left
	node.left = node.right
	node.right = temp
}

func change(node *Node) {
	node = &Node{4, nil, nil}
}
