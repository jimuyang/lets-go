package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || len(obstacleGrid) == 0 {
		return 0
	}
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}

			if j == m-1 && i == n-1 {
				obstacleGrid[i][j] = 1
				continue
			} else if j == m-1 {
				obstacleGrid[i][j] = obstacleGrid[i+1][j]
			} else if i == n-1 {
				obstacleGrid[i][j] = obstacleGrid[i][j+1]
			} else {
				obstacleGrid[i][j] = obstacleGrid[i][j+1] + obstacleGrid[i+1][j]
			}
		}
	}
	return obstacleGrid[0][0]
}

// func main() {
// 	input := [][]int{{0, 0}}
// 	uniquePathsWithObstacles(input)
// }
