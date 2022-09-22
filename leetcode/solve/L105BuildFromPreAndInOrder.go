package main

// 根据一棵树的前序遍历与中序遍历构造二叉树。

// 注意:
// 你可以假设树中没有重复的元素。

// 例如，给出
// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree1(preorder []int, inorder []int) *TreeNode {
	// 使用前序遍历的首元素在中序遍历分为2段
	return buildTree1Rec(preorder, inorder)
}

func buildTree1Rec(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}
	mid := preorder[0]
	var midIndex int
	for i, v := range inorder {
		if v == mid {
			midIndex = i
			break
		}
	}
	leftInorder := inorder[:midIndex]
	rightInorder := inorder[midIndex+1:]

	return &TreeNode{mid, buildTree1Rec(preorder[1:len(leftInorder)+1], leftInorder), buildTree1Rec(preorder[len(leftInorder)+1:], rightInorder)}
}
