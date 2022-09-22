package main

// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。

func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	// 这只能回溯吧？ 回溯个锤子
	n := len(grid)
	m := len(grid[0])

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if j == m-1 && i == n-1 {
				continue
			} else if j == m-1 {
				grid[i][j] = grid[i+1][j] + grid[i][j]
			} else if i == n-1 {
				grid[i][j] = grid[i][j+1] + grid[i][j]
			} else {
				if grid[i][j+1] < grid[i+1][j] {
					grid[i][j] = grid[i][j+1] + grid[i][j]
				} else {
					grid[i][j] = grid[i+1][j] + grid[i][j]
				}
			}
		}
	}
	return grid[0][0]
}
