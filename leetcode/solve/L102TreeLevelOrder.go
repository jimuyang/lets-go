package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	return levelOrderTravel(root, 0, result)
}

func levelOrderTravel(node *TreeNode, level int, result [][]int) [][]int {
	if node == nil {
		return result
	}
	if len(result) <= level {
		result = append(result, make([]int, 0))
	}
	result[level] = append(result[level], node.Val)

	result = levelOrderTravel(node.Left, level+1, result)
	result = levelOrderTravel(node.Right, level+1, result)
	return result
}
