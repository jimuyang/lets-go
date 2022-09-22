package main

import (
	"container/list"
)

// 给定一个二叉树，返回它的中序 遍历。

// 示例:
// 输入: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3

// 输出: [1,3,2]
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	recursiveTravel(root, &result)
	return result
}

func recursiveTravel(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	recursiveTravel(node.Left, result)
	*result = append(*result, node.Val)
	recursiveTravel(node.Right, result)
}

// 中序遍历二叉树 迭代完成 借用栈来暂存中序节点即可
func travelWithoutRec(root *TreeNode) []int {
	var stack list.List
	result := make([]int, 0)
	node := root
	arrived := false
	// 访问这个节点及子树
	for node != nil {
		// 访问left
		if node.Left != nil && !arrived {
			stack.PushBack(node)
			node = node.Left
			arrived = false
			continue
		}
		// fmt.Println(node.Val)
		result = append(result, node.Val)
		if node.Right != nil {
			node = node.Right
			arrived = false
			continue
		}
		// 访问本子树之后 栈弹出
		if stack.Len() > 0 {
			top := stack.Back()
			stack.Remove(top)
			node = top.Value.(*TreeNode)
			arrived = true
			continue
		} else {
			break
		}
	}
	return result
}
