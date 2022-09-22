package main

//给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

// 你可以对一个单词进行如下三种操作：

// 插入一个字符
// 删除一个字符
// 替换一个字符
// 示例 1:

// 输入: word1 = "horse", word2 = "ros"
// 输出: 3
// 解释:
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')

// 这个动归太巧妙了
// dp[i][j] 代表 word1 到 i 位置转换成 word2 到 j 位置需要最少步数
// 所以，
// 当 word1[i] == word2[j]，dp[i][j] = dp[i-1][j-1]；
// 当 word1[i] != word2[j]，dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
// 其中，dp[i-1][j-1] 表示替换操作，dp[i-1][j] 表示删除操作，dp[i][j-1] 表示插入操作。

// 作者：powcai
// 链接：https://leetcode-cn.com/problems/edit-distance/solution/zi-di-xiang-shang-he-zi-ding-xiang-xia-by-powcai-3/
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)

	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}

	// 第一行和第一列设值
	for i := 0; i <= len1; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= len2; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				min := dp[i-1][j-1]
				if min > dp[i][j-1] {
					min = dp[i][j-1]
				}
				if min > dp[i-1][j] {
					min = dp[i-1][j]
				}
				dp[i][j] = min + 1
			}
		}
	}
	return dp[len1][len2]
}
