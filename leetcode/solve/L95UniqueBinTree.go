package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树
func generateTrees(n int) []*TreeNode {
	// 分别以1 2 3...为根节点
	return recursiveGenerateTrees(1, n)
}

func recursiveGenerateTrees(start int, end int) []*TreeNode {
	result := make([]*TreeNode, 0)
	if start > end {
		return result
	} else if start == end {
		result = append(result, &TreeNode{start, nil, nil})
		return result
	}
	for i := start; i <= end; i++ {
		left := recursiveGenerateTrees(start, i-1)
		right := recursiveGenerateTrees(i+1, end)
		if len(left) == 0 && len(right) == 0 {
			fmt.Println("impossible")
		} else if len(left) == 0 {
			for _, r := range right {
				result = append(result, &TreeNode{i, nil, r})
			}
		} else if len(right) == 0 {
			for _, l := range left {
				result = append(result, &TreeNode{i, l, nil})
			}
		} else {
			for _, r := range right {
				for _, l := range left {
					result = append(result, &TreeNode{i, l, r})
				}
			}
		}
	}
	return result
}

// func main() {
// 	generateTrees(3)
// 	fmt.Println("end")
// }
