package main

// 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。

// 示例:
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	realHead := &ListNode{}
	realHead.Next = head

	var revHead, revTail, cutNode *ListNode
	for node, index := realHead, 0; node != nil && index <= n; index = index + 1 {
		next := node.Next
		if index == m-1 {
			cutNode = node
		} else if index == m {
			revHead = node
			revTail = node
		} else if index > m {
			node.Next = revHead
			revHead = node
		}
		if index == n {
			cutNode.Next = revHead
			revTail.Next = next
		}
		node = next
	}
	return realHead.Next
}
