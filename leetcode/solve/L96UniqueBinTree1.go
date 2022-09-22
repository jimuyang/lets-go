package main

// 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

// 动态规划
func numTrees(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		num := 0
		// 以每个数字为根节点
		for j := 1; j <= i; j++ {
			leftNum, rightNum := dp[j-1], dp[i-j]
			num += leftNum * rightNum
		}
		dp[i] = num
	}
	return dp[n]
}
