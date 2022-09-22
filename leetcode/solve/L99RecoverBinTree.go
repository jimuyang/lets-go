package main

import "fmt"

// 二叉搜索树中的两个节点被错误地交换。
// 请在不改变其结构的情况下，恢复这棵树。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	// 中序遍历途中 如何判断是哪2个节点错位了
	// 思考一个问题 一个升序的序列 交换其中2个元素的位置 可能会出现哪些情况？
	// A - S - E - B 四个连续指针向前推 其中S和E负责下降结构
	fourPtr := &FourPtr{}
	specialMidTravel(root, fourPtr)
	fourPtr.addNode(nil)
}

func specialMidTravel(node *TreeNode, fourPtr *FourPtr) {
	if node == nil {
		return
	}
	specialMidTravel(node.Left, fourPtr)
	fourPtr.addNode(node)
	specialMidTravel(node.Right, fourPtr)
}

// FourPtr 4个连续节点
type FourPtr struct {
	a, s, e, b *TreeNode
	errorNode  *TreeNode
}

func (ptr *FourPtr) addNode(node *TreeNode) bool {
	ptr.a = ptr.s
	ptr.s = ptr.e
	ptr.e = ptr.b
	ptr.b = node
	//
	return ptr.checkSE()
}

// 检查S-E是否形成下降结构 如果形成 返回问题节点
func (ptr *FourPtr) checkSE() bool {
	if ptr.s == nil || ptr.s.Val < ptr.e.Val {
		return false
	}
	var n *TreeNode
	// s e 形成了下降结构
	if ptr.a == nil || ptr.a.Val < ptr.e.Val {
		// 此时s为问题节点
		n = ptr.s
	}
	if ptr.b == nil || ptr.s.Val < ptr.b.Val {
		// 此时e为问题节点
		if n != nil {
			// s e 都是问题节点 直接交换
			swapValue(ptr.s, ptr.e)
			return true
		}
		n = ptr.e
	}
	if ptr.errorNode == nil {
		ptr.errorNode = n
		return false
	}
	swapValue(ptr.errorNode, n)
	return true
}

func swapValue(n1, n2 *TreeNode) {
	temp := n1.Val
	n1.Val = n2.Val
	n2.Val = temp
	// if np1 == nil {
	// 	// n1原来是root
	// 	n2Left, n2Right := n2.Left, n2.Right
	// 	n2.Left = n1.Left
	// 	n2.Right = n1.Right
	// 	n1.Left = n2Left
	// 	n1.Right = n2Right
	// 	if n2 == np2.Left {
	// 		np2.Left = n1
	// 	} else {
	// 		np2.Right = n1
	// 	}
	// 	return n2
	// } else if np2 == nil {
	// 	// n2原来是root
	// 	n2Left, n2Right := n2.Left, n2.Right
	// 	n2.Left = n1.Left
	// 	n2.Right = n1.Right
	// 	n1.Left = n2Left
	// 	n1.Right = n2Right
	// 	if n1 == np1.Left {
	// 		np1.Left = n2
	// 	} else {
	// 		np1.Right = n2
	// 	}
	// 	return n1
	// } else {
	// 	n2Left, n2Right := n2.Left, n2.Right
	// 	n2.Left = n1.Left
	// 	n2.Right = n1.Right
	// 	n1.Left = n2Left
	// 	n1.Right = n2Right
	// 	if n1 == np1.Left {
	// 		np1.Left = n2
	// 	} else {
	// 		np1.Right = n2
	// 	}
	// 	if n2 == np2.Left {
	// 		np2.Left = n1
	// 	} else {
	// 		np2.Right = n1
	// 	}
	// 	return nil
	// }
}

// 中序遍历后应该是一个升序数组
func midTravel(node *TreeNode) {
	if node == nil {
		return
	}
	midTravel(node.Left)
	fmt.Print(string(node.Val) + " ")
	midTravel(node.Right)
}

// func main() {
// 	root := &TreeNode{1, nil, nil}
// 	root.Left = &TreeNode{3, nil, &TreeNode{2, nil, nil}}
// 	recoverTree(root)
// 	midTravel(root)
// }
