package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree2(inorder []int, postorder []int) *TreeNode {
	return buildTree2Rec(inorder, postorder)
}

func buildTree2Rec(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{inorder[0], nil, nil}
	}
	mid := postorder[len(postorder)-1]
	var midIndex int
	for i, v := range inorder {
		if v == mid {
			midIndex = i
			break
		}
	}
	leftInorder := inorder[:midIndex]
	rightInorder := inorder[midIndex+1:]

	return &TreeNode{mid, buildTree2Rec(leftInorder, postorder[:len(leftInorder)]), buildTree2Rec(rightInorder, postorder[len(leftInorder):len(postorder)-1])}
}
