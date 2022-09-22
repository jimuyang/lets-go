package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverseLinkList(l1), reverseLinkList(l2)
	jin := 0
	fakeHead := &ListNode{}
	tail := fakeHead

	for l1 != nil || l2 != nil || jin != 0 {
		sum := jin
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			jin = 1
			sum = sum - 10
		} else {
			jin = 0
		}
		tail.Next = &ListNode{sum, nil}
		tail = tail.Next
	}
	return reverseLinkList(fakeHead.Next)
}

func reverseLinkList(head *ListNode) *ListNode {
	next := head.Next
	// 这里注意断开
	head.Next = nil
	for next != nil {
		next.Next, head, next = head, next, next.Next
	}

	return head
}

// func main() {
// 	l1 := &ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}
// 	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
// 	addTwoNumbers2(l1, l2)
// }

// 似乎如果不允许反转 就可以使用栈
