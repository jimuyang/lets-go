package main

import (
	"fmt"
	"sort"
)

// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	fmt.Println(nums)
	for index, num := range nums {
		// 简单判断和去重
		if num > 0 {
			break
		}
		if index > 0 && num == nums[index-1] {
			continue
		}
		if tr := sortTwoSum(nums[index+1:], -num); len(tr) > 0 {
			for _, item := range tr {
				result = append(result, []int{item[0], item[1], num})
			}
		}
	}
	return result
}
