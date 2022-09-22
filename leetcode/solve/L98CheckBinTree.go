package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return recIsValid(root, 0, false, 0, false)
}

func recIsValid(node *TreeNode, min int, useMin bool, max int, useMax bool) bool {
	if node == nil {
		return true
	}

	if useMin && node.Val <= min {
		return false
	}
	if useMax && node.Val >= max {
		return false
	}
	return recIsValid(node.Left, min, useMin, node.Val, true) && recIsValid(node.Right, node.Val, true, max, useMax)
}
