package main

//给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
// 示例 1:
// 输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// 输出: true

// 示例 2:
// 输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
// 输出: false
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/interleaving-string

// 动态规划试一试
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}
	dp := make([]*LinkNode, l3+1)
	mark := make([]bool, l1+1) // 标记 用于去重
	dp[0] = &LinkNode{0, 0, nil}
	for i := 0; i < l3; i++ {
		ch := s3[i]
		head := &LinkNode{0, 0, nil}
		last := head
		if dp[i] == nil {
			return false
		}
		// 遍历上次结果
		for node := dp[i]; node != nil; node = node.next {
			if node.x1 < l1 && s1[node.x1] == ch && !mark[node.x1+1] {
				mark[node.x1+1] = true
				last.next = &LinkNode{node.x1 + 1, node.x2, nil}
				last = last.next
			}
			if node.x2 < l2 && s2[node.x2] == ch && !mark[node.x1] {
				mark[node.x1] = true
				last.next = &LinkNode{node.x1, node.x2 + 1, nil}
				last = last.next
			}
		}
		dp[i+1] = head.next
		// clean mark
		for m := 0; m < l1+1; m++ {
			mark[m] = false
		}
	}
	return dp[l3] != nil
}

// LinkNode 结果node
type LinkNode struct {
	x1   int
	x2   int
	next *LinkNode
}
