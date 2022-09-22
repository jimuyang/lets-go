package main

import (
	"container/list"
)

// MinStack 最小栈
type MinStack struct {
	stack *list.List
	min   int
}

/** initialize your data structure here. */
func constructor() MinStack {
	var stack list.List
	return MinStack{&stack, 0}
}

func (this *MinStack) Push(x int) {
	if x < this.min || this.stack.Len() == 0 {
		this.min = x
	}
	this.stack.PushFront(x)
}

func (this *MinStack) Pop() {
	top := this.stack.Front()
	this.stack.Remove(top)
	if top.Value.(int) == this.min && this.stack.Len() > 0 {
		// 找到新的最小只
		min := this.stack.Front().Value.(int)
		for node := this.stack.Front(); node != nil; node = node.Next() {
			if node.Value.(int) < min {
				min = node.Value.(int)
			}
		}
		this.min = min
	}
}

func (this *MinStack) Top() int {
	return this.stack.Front().Value.(int)
}

func (this *MinStack) GetMin() int {
	if this.stack.Len() > 0 {
		return this.min
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// func main() {
// 	ms := constructor()
// 	ms.Push(1)
// 	ms.Push(2)
// 	fmt.Println(ms.Top())
// 	fmt.Println(ms.GetMin())
// 	ms.Pop()
// 	fmt.Println(ms.GetMin())
// 	fmt.Println(ms.Top())
// }
