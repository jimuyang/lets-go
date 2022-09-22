package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	head1 := &LinkedNode{1, &LinkedNode{3, &LinkedNode{5, nil}}}
	head2 := &LinkedNode{2, &LinkedNode{4, &LinkedNode{6, &LinkedNode{7, nil}}}}
	head3 := &LinkedNode{9, nil}
	input := []*LinkedNode{head1, head2, head3}
	resultHead := mergeNList(input)
	printList(resultHead)
}

func printList(node *LinkedNode) {
	var sb strings.Builder
	for ; node != nil; node = node.next {
		sb.WriteString(strconv.Itoa(node.val))
		sb.WriteString(" -> ")
	}
	fmt.Println(sb.String())
}

type LinkedNode struct {
	val  int
	next *LinkedNode
}

func mergeNList(heads []*LinkedNode) *LinkedNode {
	l := len(heads)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return heads[0]
	}
	fakeHead := &LinkedNode{}
	resultNode := fakeHead

	minHeap := newMinHeap(l)
	for i := 0; i < l; i++ {
		if heads[i] != nil {
			minHeap.add(heads[i])
		}
	}

	for minHeap.size > 0 {
		minNode, _ := minHeap.takeMin()
		resultNode.next = &LinkedNode{minNode.val, nil}
		resultNode = resultNode.next
		if minNode.next != nil {
			minHeap.add(minNode.next)
		}
	}
	return fakeHead.next
}

// MinHeap 最小堆
type MinHeap struct {
	heap     []*LinkedNode
	size     int
	capacity int
}

func newMinHeap(cap int) *MinHeap {
	heap := make([]*LinkedNode, cap+1)
	return &MinHeap{heap, 0, cap}
}

func (me *MinHeap) min() (*LinkedNode, error) {
	if me.size <= 0 {
		return nil, fmt.Errorf("heap is empty")
	}
	return me.heap[1], nil
}

func (me *MinHeap) takeMin() (*LinkedNode, error) {
	if me.size <= 0 {
		return nil, fmt.Errorf("heap is empty")
	}
	min := me.heap[1]
	me.heap[1] = me.heap[me.size]
	me.size--
	me.heapify(1)
	return min, nil
}

func (me *MinHeap) add(val *LinkedNode) error {
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
	for i > 1 && me.heap[i/2].val > me.heap[i].val {
		me.heap[i/2], me.heap[i] = me.heap[i], me.heap[i/2]
		i = i / 2
	}
}

// 自顶而下
func (me *MinHeap) heapify(i int) {
	min := i
	// parent和left的较小值
	if 2*i <= me.size && me.heap[2*i].val < me.heap[i].val {
		min = 2 * i
	}
	if 2*i+1 <= me.size && me.heap[2*i+1].val < me.heap[min].val {
		min = 2*i + 1
	}
	if min != i {
		// 需要调整
		me.heap[min], me.heap[i] = me.heap[i], me.heap[min]
		me.heapify(min)
	}
}

// func minHeap(nums []int) []int {
// 	// copy
// 	heap, heapSize := make([]*LinkedNode, len(nums)+1), len(nums)

// 	for i := 1; i < len(nums)+1; i++ {
// 		heap[i] = nums[i-1]
// 	}
// 	// 对每一个非页节点做调整
// 	for i := heapSize / 2; i >= 1; i-- {
// 		minHeapify(heap, i)
// 	}
// 	fmt.Println(heap)
// 	return heap
// }

func minHeapify(heap []*LinkedNode, i int) {
	min := i
	// parent和left的较小值
	if 2*i <= len(heap) && heap[2*i].val < heap[i].val {
		min = 2 * i
	}
	if 2*i+1 <= len(heap) && heap[2*i+1].val < heap[min].val {
		min = 2*i + 1
	}
	if min != i {
		// 需要调整
		heap[min], heap[i] = heap[i], heap[min]
		minHeapify(heap, min)
	}
}
