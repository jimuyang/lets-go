package main

// 给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
// 对每行做一个切面 问题转化为在柱状图中寻找最大矩形
func maximalRectangle(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 {
		return 0
	}
	result := 0
	heights := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}
		r := largestRectangleArea2(heights)
		if r > result {
			result = r
		}
	}
	return result
}

// func main() {
// 	matrix := [][]int{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}
// 	fmt.Println(maximalRectangle(matrix))
// }
