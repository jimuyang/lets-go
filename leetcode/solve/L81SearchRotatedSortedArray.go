package main

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
// ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
// 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
// 示例 1:
// 输入: nums = [2,5,6,0,0,1,2], target = 0
// 输出: true
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func search(nums []int, target int) bool {
	if len(nums) < 1 {
		return false
	}
	if len(nums) == 1 {
		return target == nums[0]
	}
	if len(nums) == 2 {
		return target == nums[0] || target == nums[1]
	}
	// 长度大于2时 使用二分搜索
	midI := len(nums) / 2
	head, mid, tail := nums[0], nums[len(nums)/2], nums[len(nums)-1]
	if head == target || tail == target || mid == target {
		return true
	}

	// 简化head==tail的场景
	if head == tail {
		for i := 1; i < len(nums); i++ {
			if nums[i] != tail {
				return search(nums[i:], target)
			}
		}
		return false
	}

	if head > tail {
		// 这一段内有旋转
		if head <= mid {
			// 左边正常顺序
			if target > mid || target < tail {
				return search(nums[midI:], target)
			} else {
				return search(nums[:midI], target)
			}
		} else {
			// 右边正常顺序
			if target < mid || target > head {
				return search(nums[:midI], target)
			} else {
				return search(nums[midI:], target)
			}
		}
	} else {
		// 这一段内无旋转
		if target > mid {
			return search(nums[midI:], target)
		} else {
			return search(nums[:midI], target)
		}
	}
}

// func main() {
// 	fmt.Println(search([]int{1, 3, 1, 1}, 3))
// }
