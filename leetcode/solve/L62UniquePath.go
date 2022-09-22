package main

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
// How many possible unique paths are there?

// 先来个递归版本 超时。。。
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}

// 动归试试吧
func uniquePaths1(m int, n int) int {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}

	// 首行 首列置1
	for i := 0; i < m; i++ {
		arr[i][0] = 1
	}
	for i := 0; i < n; i++ {
		arr[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			arr[j][i] = arr[j][i-1] + arr[j-1][i]
		}
	}
	// fmt.Println(arr)
	return arr[m-1][n-1]
}

// func main() {
// 	uniquePaths1(7, 3)
// }
