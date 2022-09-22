package main

// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var b *ListNode
	for node := head; node != nil; node = node.Next {
		if b == nil {
			b = node
			continue
		}
		if node.Val == b.Val {
			b.Next = node.Next
		} else {
			b = node
		}
	}
	return head
}
