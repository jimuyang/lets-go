package main

// 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
// type ListNode struct {
// Val  int
// Next *ListNode
// }
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 太累了 加个假head
	fakeHead := ListNode{head.Val - 1, head}
	bb := &fakeHead
	var b *ListNode
	dupB := false
	for node := head; node != nil; node = node.Next {
		if b == nil {
			b = node
			continue
		}
		if node.Val == b.Val {
			dupB = true
			b.Next = node.Next
		} else {
			if dupB {
				bb.Next = node
			} else {
				bb = b
			}
			b = node
			dupB = false
		}
	}
	if dupB {
		bb.Next = nil
	}

	return fakeHead.Next
}
