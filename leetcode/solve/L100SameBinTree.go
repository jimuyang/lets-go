package main

// 给定两个二叉树，编写一个函数来检验它们是否相同。
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	pArr := make([]*TreeNode, 2)
	pArr[0] = nil
	pArr[1] = p
	qArr := make([]*TreeNode, 2)
	qArr[0] = nil
	qArr[1] = q

	// 树的广度优先遍历
	notEmpty := true
	for i := 0; notEmpty; i++ {
		notEmpty = false
		for j := 1 << i; j < 1<<(i+1); j++ {
			if (pArr[j] == nil && qArr[j] == nil) ||
				(pArr[j] != nil && qArr[j] != nil && pArr[j].Val == qArr[j].Val) {
				// 相等
			} else {
				return false
			}
			if pArr[j] != nil {
				notEmpty = notEmpty || pArr[j].Left != nil
				pArr = append(pArr, pArr[j].Left)
				notEmpty = notEmpty || pArr[j].Right != nil
				pArr = append(pArr, pArr[j].Right)
			} else {
				pArr = append(pArr, nil)
				pArr = append(pArr, nil)
			}

			if qArr[j] != nil {
				qArr = append(qArr, qArr[j].Left)
				qArr = append(qArr, qArr[j].Right)
			} else {
				qArr = append(qArr, nil)
				qArr = append(qArr, nil)
			}
		}
	}
	return true
}
