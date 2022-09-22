package main

import "fmt"

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

// 递归试试呢 4-2 = 3-2 + 3-1
var result [][]int

func combine(n int, k int) [][]int {
	up, down := 1, 1
	for i := 0; i < k; i++ {
		down *= i + 1
		up *= n - i
	}
	count := up / down
	fmt.Println("count", count)

	result = make([][]int, count)
	for i := 0; i < count; i++ {
		// result = append(result, make([]int, k))
		result[i] = make([]int, k)
	}
	fmt.Println(result)
	reccombine(n, k, 0)
	return result
}

func reccombine(n int, k int, offset int) (left int, right int) {
	if k == 1 {
		for i := 0; i < n; i++ {
			// result[offset+i] = append(result[offset+i], i+1)
			result[offset+i][0] = i + 1
		}
		return offset, offset + n
	}

	if n == k {
		for i := 0; i < n; i++ {
			result[offset][i] = i + 1
		}
		return offset, offset + 1
	}

	l1, r1 := reccombine(n-1, k, offset)
	l2, r2 := reccombine(n-1, k-1, r1)
	for _, s := range result[l2:r2] {
		s[k-1] = n
	}
	return l1, r2

}
