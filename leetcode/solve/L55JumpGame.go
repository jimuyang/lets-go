package main

/*Given an array of non-negative integers, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Determine if you are able to reach the last index.

Example 1:
Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index. */
func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	return jumpRecursive(nums, len(nums)-1)
}

//jumpRecursive 是一种递归方式 效率很低
func jumpRecursive(num []int, target int) bool {
	if target == 0 {
		return true
	}
	canArrive := false
	for i := 0; i < target; i++ {
		// 能否路过i到达target
		canArrive = canArrive || (num[i]+i >= target && jumpRecursive(num, i))
	}
	return canArrive
}

// jumpAndMark 通过开辟等长的数组来标记哪些点可达
func jumpAndMark(nums []int) bool {
	numsLen := len(nums)
	if numsLen < 2 {
		return true
	}
	if nums[0] < 1 {
		return false
	}
	cursor := nums[0]
	for i := 0; i <= cursor; i++ {
		if nums[i]+i > cursor {
			cursor = nums[i] + i
		}
		if cursor >= numsLen-1 {
			return true
		}
	}
	return false
}

// func main() {
// }
