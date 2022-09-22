package main

// 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	// 转数组来实现是ok的方案
	nodes := make([]*ListNode, 0)
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	// 递归二分形成树
	return buildTree(nodes)
}

func buildTree(nodes []*ListNode) *TreeNode {
	l := len(nodes)
	if l < 1 {
		return nil
	}
	if l == 1 {
		return &TreeNode{nodes[0].Val, nil, nil}
	}
	mid := l / 2
	return &TreeNode{nodes[mid].Val, buildTree(nodes[:mid]), buildTree(nodes[mid+1:])}
}
