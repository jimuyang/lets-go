package main

// 给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
// 借用第一行和第一列来存储该行该列是否为0
func setZeroes(matrix [][]int) {
	width := len(matrix)
	if width < 1 {
		return
	}
	length := len(matrix[0])
	firstCol := false
	firstRow := false
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					firstCol = true
				}
				if j == 0 {
					firstRow = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < width; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < length; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < length; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < width; i++ {
				matrix[i][j] = 0
			}
		}
	}

	if firstCol {
		for j := 0; j < length; j++ {
			matrix[0][j] = 0
		}
	}
	if firstRow {

		for i := 0; i < width; i++ {
			matrix[i][0] = 0
		}
	}
}
