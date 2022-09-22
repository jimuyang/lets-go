package main

import (
	"container/list"
	"fmt"
)

// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	var stack list.List
	stack.PushFront(root)
	stackSum := root.Val
	stackSum = findPathToLeaf(&stack, stackSum)
	if stackSum == sum {
		return true
	}
	for stack.Len() > 0 {
		pop := stack.Front().Value.(*TreeNode)
		stack.Remove(stack.Front())
		stackSum -= pop.Val
		if stack.Len() == 0 {
			break
		}
		top := stack.Front().Value.(*TreeNode)
		// fmt.Println("pop:", pop.Val, " top:", top.Val)
		if pop == top.Left && top.Right != nil {
			stack.PushFront(top.Right)
			stackSum += top.Right.Val
			stackSum = findPathToLeaf(&stack, stackSum)
			if stackSum == sum {
				return true
			}
		}
	}
	return false
}

func findPathToLeaf(stack *list.List, stackSum int) int {
	if stack.Len() > 0 {
		node := stack.Front().Value.(*TreeNode)
		for node.Left != nil || node.Right != nil {
			if node.Left != nil {
				node = node.Left
			} else if node.Right != nil {
				node = node.Right
			}
			stackSum += node.Val
			stack.PushFront(node)
		}
		// printStack(stack)
		return stackSum
	}
	panic("stack is empty")
}

func printStack(stack *list.List) {
	for n := stack.Back(); n != nil; n = n.Prev() {
		fmt.Print(n.Value.(*TreeNode).Val, "->")
	}
	fmt.Println()
}
