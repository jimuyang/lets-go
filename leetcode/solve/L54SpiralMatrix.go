package main

import (
	"fmt"
)

/**
 * Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order
 * Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
*/
func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return make([]int, 0)
	}
	width := len(matrix[0])
	height := len(matrix)
	// resultLen := width * height
	var result []int

	var flag int
	// 层数
	for n := 0; n < width && n < height; n++ {
		// x := n
		// y := n
		fmt.Println(n)

		flag = 0
		// 向右
		for i := n; i < width-n; i++ {
			result = append(result, matrix[n][i])
			flag++
		}
		if flag == 0 {
			break
		}

		flag = 0
		// 向下
		for i := n + 1; i < height-n; i++ {
			result = append(result, matrix[i][width-n-1])
			flag++
		}
		if flag == 0 {
			break
		}

		flag = 0
		// 向左
		for i := width - n - 2; i >= n; i-- {
			result = append(result, matrix[height-n-1][i])
			flag++
		}
		if flag == 0 {
			break
		}

		flag = 0
		// 向上
		for i := height - n - 2; i > n; i-- {
			result = append(result, matrix[i][n])
			flag++
		}
		if flag == 0 {
			break
		}
	}
	return result
}

// func main() {
// 	matrix := [][]int{
// 		{1, 2, 3, 4},
// 		{5, 6, 7, 8},
// 		{9, 10, 11, 12},
// 	}
// 	fmt.Println(spiralOrder(matrix))
// }
