package main

// 给定一颗二叉树 判断是否 高度平衡
// 高度平衡 height-balanced 每一个节点的左右子树高度相差不超过1
func isBalanced(root *TreeNode) bool {
	ok, _ := isBalancedWithHeight(root)
	return ok
}

func isBalancedWithHeight(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	leftOk, leftH := isBalancedWithHeight(root.Left)
	if !leftOk {
		return false, 0
	}
	rightOk, rightH := isBalancedWithHeight(root.Right)
	if !rightOk {
		return false, 0
	}
	if leftH > rightH {
		return leftH-rightH <= 1, leftH + 1
	} else {
		return rightH-leftH <= 1, rightH + 1
	}
}
