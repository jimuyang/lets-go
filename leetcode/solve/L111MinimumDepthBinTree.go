package main

// 给定一颗二叉树 找出最小深度

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		left, right := minDepth(root.Left), minDepth(root.Right)
		if left < right {
			return left + 1
		} else {
			return right + 1
		}
	}
}
