package main

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

func twoSum(nums []int, target int) []int {
	// 借用一个map结构来实现单趟 map[缺值]index
	mp := make(map[int]int)
	for index, num := range nums {
		if value, exist := mp[num]; exist {
			return []int{value, index}
		}
		mp[target-num] = index
	}
	return []int{}
}

// 已知nums升序情况下的twoSum 从两侧往中间夹逼
func sortTwoSum(nums []int, target int) [][]int {
	result := make([][]int, 0)

	lo, hi := 0, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum == target {
			result = append(result, []int{nums[lo], nums[hi]})
			for lo < hi && nums[lo] == nums[lo+1] {
				lo++
			}
			for lo < hi && nums[hi] == nums[hi-1] {
				hi--
			}
			lo++
			hi--
		} else if sum < target {
			lo++
		} else {
			hi--
		}
	}
	return result
}
