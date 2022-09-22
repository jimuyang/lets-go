package main

//给定一个二叉树，检查它是否是镜像对称的。
//如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 试试看 对于根节点的2个子树 以对称的顺序遍历
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricLR(root.Left, root.Right)
}

func isSymmetricLR(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil && left.Val == right.Val {
		// left == right 检查left.left == right.right && left.right == right.left
		return isSymmetricLR(left.Left, right.Right) && isSymmetricLR(left.Right, right.Left)
	} else {
		return false
	}
}
