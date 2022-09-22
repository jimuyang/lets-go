package main

// ListNode 链表节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	node := head
	// 先计算链表长度
	listLen := 1
	for ; node.Next != nil; listLen++ {
		node = node.Next
	}
	lastNode := node
	// fmt.Println(listLen)

	if k >= listLen {
		k = k % listLen
	}
	if k == 0 {
		return head
	}

	breakPoint := listLen - k

	// 找到breakPoint
	node = head
	for i := 1; i < breakPoint; i++ {
		node = node.Next
	}
	// fmt.Println(node.Val)
	lastNode.Next = head
	head = node.Next
	node.Next = nil
	// breakNode := node
	// node = breakNode.Next
	// breakNode.Next = nil
	// for node != nil {
	// 	nextNode := node.Next
	// 	node.Next = head
	// 	head = node
	// 	node = nextNode
	// }
	return head
}

// func main() {
// 	head := new(ListNode)
// 	head.Val = 1

// 	node1 := new(ListNode)
// 	node1.Val = 2
// 	head.Next = node1

// 	node2 := new(ListNode)
// 	node2.Val = 3
// 	node1.Next = node2

// 	node3 := new(ListNode)
// 	node3.Val = 4
// 	node2.Next = node3

// 	node4 := new(ListNode)
// 	node4.Val = 5
// 	node3.Next = node4

// 	rotateRight(head, 2)
// }
