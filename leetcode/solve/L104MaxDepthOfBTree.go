package main

import "container/list"

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	ld, rd := maxDepth(node.Left), maxDepth(node.Right)
	if ld < rd {
		return rd + 1
	}
	return ld + 1
}

// 非递归版本
func maxDepth1(node *TreeNode) int {
	var l list.List
	tag := &TreeNode{}
	if node == nil {
		return 0
	}
	l.PushBack(node)
	l.PushBack(tag)
	height := 0
	for listNode := l.Front(); listNode != nil; listNode = listNode.Next() {
		node := listNode.Value.(*TreeNode)
		if node == tag {
			height++
			if listNode.Next() == nil {
				break
			}
			l.PushBack(tag)
			continue
		}
		if node == nil {
			panic("node == nil")
		}
		if node.Left != nil {
			l.PushBack(node.Left)
		}
		if node.Right != nil {
			l.PushBack(node.Right)
		}
	}
	return height
}
