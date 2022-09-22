package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{})
	if len(nums) == 0 {
		return result
	}
	sort.Ints(nums)
	usedIndex := 0
	for i := 0; i < len(nums); i++ {
		same := false
		if i > 0 {
			same = nums[i] == nums[i-1]
		}

		if !same {
			usedIndex = len(result)
			for j := 0; j < usedIndex; j++ {
				l := len(result[j]) + 1
				t := make([]int, l)
				t[l-1] = nums[i]
				copy(t[:l-1], result[j])
				result = append(result, t)
			}
		} else {
			// 如果和上一个值一样
			lastUsedIndex := usedIndex
			usedIndex = len(result)
			for j := lastUsedIndex; j < usedIndex; j++ {
				l := len(result[j]) + 1
				t := make([]int, l)
				t[l-1] = nums[i]
				copy(t[:l-1], result[j])
				result = append(result, t)
			}
		}
	}
	return result
}
