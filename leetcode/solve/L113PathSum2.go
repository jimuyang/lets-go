package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
// Note: A leaf is a node with no children.

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	path := &BinTreePath{path: list.New()}
	path.push(root)
	path.tryExpandPath()
	if path.sum == sum {
		res = append(res, path.travel())
	}

	for err := path.next(); err == nil; err = path.next() {
		if path.sum == sum {
			res = append(res, path.travel())
		}
	}
	return res
}

type BinTreePath struct {
	path *list.List
	sum  int
}

func (p *BinTreePath) push(node *TreeNode) {
	p.path.PushFront(node)
	p.sum += node.Val
}

func (p *BinTreePath) pop() *TreeNode {
	removed := p.path.Remove(p.path.Front()).(*TreeNode)
	p.sum -= removed.Val
	return removed
}

func (p *BinTreePath) next() error {
	path := p.path
	if path.Len() <= 1 {
		return fmt.Errorf("no next path")
	}
	for path.Len() > 0 {
		removed := p.pop()
		if path.Len() == 0 {
			return fmt.Errorf("no next path")
		}
		node := path.Front().Value.(*TreeNode)
		if node.Right == removed || node.Right == nil {
			continue
		} else {
			p.push(node.Right)
			p.tryExpandPath()
			break
		}
	}
	return nil
}

// 延长path
func (p *BinTreePath) tryExpandPath() {
	path := p.path
	node := path.Front().Value.(*TreeNode)
	for node.Left != nil || node.Right != nil {
		var expanded *TreeNode
		if node.Left != nil {
			expanded = node.Left
		} else {
			expanded = node.Right
		}
		p.push(expanded)
		node = expanded
	}
	p.show()
}

func (p *BinTreePath) show() {
	stack := p.path
	res := make([]string, 0, stack.Len())
	for node := stack.Back(); node != nil; node = node.Prev() {
		val := node.Value.(*TreeNode).Val
		res = append(res, strconv.Itoa(val))
	}
	fmt.Println(strings.Join(res, " -> "), p.sum)
}

func (p *BinTreePath) travel() []int {
	stack := p.path
	res := make([]int, 0, stack.Len())
	for node := stack.Back(); node != nil; node = node.Prev() {
		res = append(res, node.Value.(*TreeNode).Val)
	}
	return res
}
