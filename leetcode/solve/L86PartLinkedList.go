package main

// 链表分割

// 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
// 你应当保留两个分区中每个节点的初始相对位置。

// 示例:
// 输入: head = 1->4->3->2->5->2, x = 3
// 输出: 1->2->2->4->3->5
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var left, right *ListNode
	var lh, rh *ListNode
	for node := head; node != nil; node = node.Next {
		switch {
		case node.Val >= x:
			if right == nil {
				right = node
				rh = node
			} else {
				right.Next = node
				right = node
			}
		case node.Val < x:
			if left == nil {
				left = node
				lh = node
			} else {
				left.Next = node
				left = node
			}
		}
	}
	if right != nil {
		right.Next = nil
	}
	if left != nil {
		left.Next = rh
	}

	var result *ListNode
	if lh != nil {
		result = lh
	} else {
		result = rh
	}
	return result
}

// func main() {
// 	head := &ListNode{1, nil}
// 	head.Next = &ListNode{4, nil}
// 	head.Next.Next = &ListNode{3, nil}
// 	head.Next.Next.Next = &ListNode{2, nil}
// 	head.Next.Next.Next.Next = &ListNode{5, nil}
// 	head.Next.Next.Next.Next.Next = &ListNode{2, nil}

// 	r := partition(head, 3)
// 	fmt.Println(r)
// }
