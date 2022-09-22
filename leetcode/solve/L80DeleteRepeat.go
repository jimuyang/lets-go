package main

// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

// 示例 1:
// 给定 nums = [1,1,1,2,2,3],
// 函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。
// 你不需要考虑数组中超出新长度后面的元素。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	wi := 0 // 写index
	count := 0
	for ri := 0; ri < len(nums); ri++ {
		if ri == 0 {
			wi++
			count = 1
			continue
		}

		if nums[ri] == nums[wi-1] {
			// 重复的数字
			if count < 2 {
				// 还能放
				count++
				nums[wi] = nums[ri]
				wi++
			}
		} else {
			// 不重复的数字
			count = 1
			nums[wi] = nums[ri]
			wi++
		}
	}
	return wi
}
