package main

// 给出二叉搜索树的根节点，该二叉树的节点值各不相同，修改二叉树，使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

// 提醒一下，二叉搜索树满足下列约束条件：

// 节点的左子树仅包含键小于节点键的节点。
// 节点的右子树仅包含键大于节点键的节点。
// 左右子树也必须是二叉搜索树。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func bstToGst(root *TreeNode) *TreeNode {
	rolLoop(root, 0)
	return root
}

// func test1(node *TreeNode, s int) int {
// 	sum := test1(node.Right, s)
// 	node.Val = node.Val + sum
// 	return test1(node.Left, node.Val)
// }

// right o left
func rolLoop(node *TreeNode, sum int) int {
	if node == nil {
		return sum
	}
	rsum := rolLoop(node.Right, sum)
	node.Val = node.Val + rsum
	return rolLoop(node.Left, node.Val)
}
