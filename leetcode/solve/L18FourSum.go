package main

import "sort"

// 四数之和
// 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
// 注意：
// 答案中不可以包含重复的四元组。
func fourSum(nums []int, target int) [][]int {
	return NSum(nums, 4, target)
}

// NSum 通常的 从nums中找到n个元素 使得这个n个元素的和为target 不可重复
func NSum(nums []int, n int, target int) [][]int {
	// 排序
	sort.Ints(nums)
	return recNSum(nums, n, target)
}

// 递归版本 很容易实现 (前提nums有序)
func recNSum(nums []int, n int, target int) [][]int {
	result := make([][]int, 0)
	// n == 1 时没有意义
	if n == 2 {
		return sortTwoSum(nums, target)
	}
	// n > 2 时遍历
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		if n1Result := recNSum(nums[i+1:], n-1, target-num); len(n1Result) > 0 {
			// n-1的结果
			for _, n1Item := range n1Result {
				r := make([]int, n)
				for i, v := range n1Item {
					r[i] = v
				}
				r[n-1] = num
				result = append(result, r)
			}
		}
	}
	return result
}
