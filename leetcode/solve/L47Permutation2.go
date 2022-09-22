package main

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	result := make([][]int, 0)
	result = append(result, append(make([]int, 0), nums...))

	for {
		// 从尾部开始寻找降序结构
		i := len(nums) - 1
		for ; i >= 1 && nums[i] <= nums[i-1]; i-- {
		}
		if i == 0 {
			// 全降序
			return result
		}
		ti := i - 1
		// 找到第一个比它大的数
		for i < len(nums) && nums[i] > nums[ti] {
			i++
		}
		// 交换
		nums[i-1], nums[ti] = nums[ti], nums[i-1]
		if ti != len(nums)-1 {
			reverseArray(nums[ti+1:])
		}
		result = append(result, append(make([]int, 0), nums...))
	}
}

// 反序
func reverseArray(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// func main() {
// 	fmt.Println(permuteUnique([]int{1, 1, 2, 3}))
// }
