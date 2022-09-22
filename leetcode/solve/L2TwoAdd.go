package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultHead, resultNode *ListNode
	jin := 0
	// 三者不能全为空
	for !(jin == 0 && l1 == nil && l2 == nil) {
		v1, v2, v3 := 0, 0, jin
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		v := v1 + v2 + v3
		if v >= 10 {
			v = v - 10
			jin = 1
		} else {
			jin = 0
		}
		newNode := &ListNode{v, nil}
		if resultHead == nil {
			resultHead = newNode
			resultNode = newNode
		} else {
			resultNode.Next = newNode
			resultNode = newNode
		}
	}
	return resultHead
}

// ListNode 连标
type ListNode struct {
	Val  int
	Next *ListNode
}
