package main

import (
	"container/list"
	"strconv"
	"strings"
)

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [5,4,8,11,null,13,4,7,2,null,null,5,1]
func buildTreeByHeap(heap string) *TreeNode {
	trimed := strings.TrimSpace(strings.Trim(strings.Trim(heap, "["), "]"))
	if trimed == "" {
		return nil
	}
	arr := strings.Split(trimed, ",")
	nodes := make([]*TreeNode, len(arr))

	for i, val := range arr {
		if val == "null" {
			continue
		}
		node := &TreeNode{Val: atoi(arr[i])}
		nodes[i] = node
		if i == 0 {
			continue
		}
		parent := (i+1)/2 - 1
		isLeftChild := i%2 == 1
		if isLeftChild {
			nodes[parent].Left = node
		} else {
			nodes[parent].Right = node
		}
	}
	return nodes[0]
}

func buildTreeByBreadthFirst(input string) *TreeNode {
	trimed := strings.TrimSpace(strings.Trim(strings.Trim(input, "["), "]"))
	if trimed == "" {
		return nil
	}
	arr := strings.Split(trimed, ",")
	l := list.New()
	root := &TreeNode{Val: atoi(arr[0])}
	l.PushBack(root)

	for i := 1; i < len(arr); {
		left, right := arr[i], arr[i+1]
		i += 2
		parent := l.Remove(l.Front()).(*TreeNode)
		if left != "null" {
			leftNode := &TreeNode{Val: atoi(left)}
			parent.Left = leftNode
			l.PushBack(leftNode)
		}
		if right != "null" {
			rightNode := &TreeNode{Val: atoi(right)}
			parent.Right = rightNode
			l.PushBack(rightNode)
		}
	}
	return root
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
