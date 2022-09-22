package main

// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	result := make([][]int, 0)
	result = append(result, []int{})
	result = append(result, []int{nums[len(nums)-1]})

	for i := len(nums) - 2; i >= 0; i-- {
		l := len(result)
		for j := 0; j < l; j++ {
			t := make([]int, len(result[j])+1)
			t[0] = nums[i]
			copy(t[1:], result[j])
			result = append(result, t)
			// result = append(result, result[j])
			// result[j] = append(result[j], nums[i])
		}
	}
	return result
}

// func main() {
// 	fmt.Println(subsets([]int{1, 2, 3}))
// }

// func recSubsets(nums []int) [][]int {
// 	if len(nums) == 1 {
// 		result := make([][]int, 1)
// 		result[0] = []int{nums[0]}
// 		return result
// 	}

// 	recSubsets(nums[:len(nums)-1])

// }
