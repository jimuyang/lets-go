package main

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	left := true
	var nodeList list.List
	nodeList.PushBack(root)
	for ; nodeList.Len() > 0; left = !left {
		arr, tempList := travelNodeList(nodeList, left)
		if len(arr) > 0 {
			result = append(result, arr)
		}
		nodeList = tempList
	}
	return result
}

func travelNodeList(l list.List, left bool) ([]int, list.List) {
	arr := make([]int, 0)
	var resultList list.List
	for ln := l.Front(); ln != nil; ln = ln.Next() {
		node := ln.Value.(*TreeNode)
		if node != nil {
			arr = append(arr, node.Val)
			if left {
				resultList.PushFront(node.Left)
				resultList.PushFront(node.Right)
			} else {
				resultList.PushFront(node.Right)
				resultList.PushFront(node.Left)
			}
		}
	}
	return arr, resultList
}
