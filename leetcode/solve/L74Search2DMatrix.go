package main

import (
	"fmt"
)

// 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
func searchMatrix(matrix [][]int, target int) bool {
	ml := len(matrix)
	if ml == 0 || len(matrix[0]) == 0 {
		return false
	}
	// 先确定是在哪一行？
	li, ri := 0, ml-1
	left, right := matrix[li][0], matrix[ri][0]

	ti := -1
	if ml == 1 {
		// 就一行
		ti = 0
	} else if target == left || target == right {
		return true
	} else if target < left {
		return false
	} else if target > right {
		// 认为在最后一行
		ti = ml - 1
	} else {
		// 二分查找确定区间
		for {
			midi := (li + ri) / 2
			if matrix[midi][0] == target {
				return true
			}
			if midi == li {
				ti = midi
				break
			}
			if matrix[midi][0] < target {
				li = midi
			} else {
				ri = midi
			}
		}
	}

	fmt.Println("ti:", ti)
	// 找到目标行之后 再二分查找是否存在
	l := len(matrix[ti])
	li, ri = 0, l-1
	if matrix[ti][l-1] == target {
		return true
	} else if matrix[ti][l-1] < target {
		return false
	} else {
		for {
			midi := (li + ri) / 2
			if matrix[ti][midi] == target {
				return true
			}
			if midi == li {
				return false
			}
			if matrix[ti][midi] < target {
				li = midi
			} else {
				ri = midi
			}
		}
	}
}
