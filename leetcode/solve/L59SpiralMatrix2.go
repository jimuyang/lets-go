package main

import (
	"fmt"
)

// Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.
// Example:
// Input: 3
// Output:
// [
//  [ 1, 2, 3 ],
//  [ 8, 9, 4 ],
//  [ 7, 6, 5 ]
// ]

func generateMatrix(n int) [][]int {
	matrix := [][]int{}
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}
	fmt.Println(matrix)

	flag := 0
	num := 1
	// 层数
	for j := 0; j < n; j++ {

		flag = 0
		// 向右
		for i := j; i < n-j; i++ {
			matrix[j][i] = num
			num++
			flag++
		}
		if flag == 0 {
			break
		}

		flag = 0
		// 向下
		for i := j + 1; i < n-j; i++ {
			matrix[i][n-j-1] = num
			num++
			flag++
		}
		if flag == 0 {
			break
		}

		flag = 0
		// 向左
		for i := n - j - 2; i >= j; i-- {
			matrix[n-j-1][i] = num
			num++
			flag++
		}
		if flag == 0 {
			break
		}

		flag = 0
		// 向上
		for i := n - j - 2; i > j; i-- {
			matrix[i][j] = num
			num++
			flag++
		}
		if flag == 0 {
			break
		}
	}
	return matrix
}
