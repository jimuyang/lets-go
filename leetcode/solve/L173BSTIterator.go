package main

import (
	"container/list"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * 本质上是中序遍历 当然是用栈存储状态的中序遍历
 */
type BSTIterator struct {
	stack *list.List
}

func Constructor1(root *TreeNode) BSTIterator {
	// 左孩子一路入栈
	var stack list.List
	for node := root; node != nil; node = node.Left {
		stack.PushFront(node)
	}
	return BSTIterator{&stack}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	// 栈顶就是next
	if this.stack.Len() > 0 {
		top := this.stack.Front()
		this.stack.Remove(top)
		node := top.Value.(*TreeNode)
		if node.Right != nil {
			for n := node.Right; n != nil; n = n.Left {
				this.stack.PushFront(n)
			}
		}
		return node.Val
	}
	return -1
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.stack.Len() > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func midOrderTravel(root *TreeNode) []int {
	result := make([]int, 0)
	var stack list.List
	// 左孩子一路入栈
	for node := root; node != nil; node = node.Left {
		stack.PushFront(node)
	}
	for stack.Len() > 0 {
		// 开始退栈
		top := stack.Front().Value.(*TreeNode)
		stack.Remove(stack.Front())
		// 访问该节点
		// fmt.Println(top.Val)
		result = append(result, top.Val)
		// 处理右孩子
		if top.Right != nil {
			// 左孩子一路入栈
			for node := top.Right; node != nil; node = node.Left {
				stack.PushFront(node)
			}
		}
	}
	return result
}
