package main

// 一条包含字母 A-Z 的消息通过以下方式进行了编码：

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。

// 示例 1:

// 输入: "12"
// 输出: 2
// 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。

func numDecodings(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	// 由短到长 动态规划 很像fibnacci 需要单独处理0
	dp := make([]int, l+1)
	dp[0] = 1
	for i := 0; i < l; i++ {
		if i == 0 && s[0] == '0' {
			return 0
		}
		if i > 0 && s[i] == '0' && (s[i-1] > '2' || s[i-1] == '0') {
			return 0
		}
		if i > 0 && ((s[i-1] == '2' && s[i] <= '6') || s[i-1] == '1') {
			if s[i] == '0' {
				dp[i+1] = dp[i-1]
			} else {
				dp[i+1] = dp[i] + dp[i-1]
			}
		} else {
			dp[i+1] = dp[i]
		}
	}
	return dp[l]
}

// func main() {
// 	fmt.Println(numDecodings("101"))
// }
