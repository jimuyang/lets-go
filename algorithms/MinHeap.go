package main

import "fmt"

// MinHeap 最小堆
type MinHeap struct {
	heap     []Interface
	size     int
	capacity int
}

type Interface interface {
	getValue() int
}

func newMinHeap(cap int) *MinHeap {
	heap := make([]Interface, cap+1)
	return &MinHeap{heap, 0, cap}
}

func (me *MinHeap) min() (Interface, error) {
	if me.size <= 0 {
		return nil, fmt.Errorf("heap is empty")
	}
	return me.heap[1], nil
}

func (me *MinHeap) takeMin() (Interface, error) {
	if me.size <= 0 {
		return nil, fmt.Errorf("heap is empty")
	}
	min := me.heap[1]
	me.heap[1] = me.heap[me.size]
	me.size--
	me.heapify(1)
	return min, nil
}

func (me *MinHeap) add(val Interface) error {
	if me.size >= me.capacity {
		return fmt.Errorf("heap is full")
	}
	me.size++
	me.heap[me.size] = val
	me.bubble(me.size)
	return nil
}

// 自底而上 冒泡
func (me *MinHeap) bubble(i int) {
	for i > 1 && me.heap[i/2].getValue() > me.heap[i].getValue() {
		me.heap[i/2], me.heap[i] = me.heap[i], me.heap[i/2]
		i = i / 2
	}
}

// 自顶而下
func (me *MinHeap) heapify(i int) {
	min := i
	// parent和left的较小值
	if 2*i <= me.size && me.heap[2*i].getValue() < me.heap[i].getValue() {
		min = 2 * i
	}
	if 2*i+1 <= me.size && me.heap[2*i+1].getValue() < me.heap[min].getValue() {
		min = 2*i + 1
	}
	if min != i {
		// 需要调整
		me.heap[min], me.heap[i] = me.heap[i], me.heap[min]
		me.heapify(min)
	}
}

type TestType struct {
	val int
}

func (tt TestType) getValue() int {
	return tt.val
}

// func main() {
// 	minHeap := newMinHeap(10)
// 	minHeap.add(TestType{})
// }
