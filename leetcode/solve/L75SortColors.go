package main

// 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sort-colors
func sortColors(nums []int) {
	l := len(nums)
	if nums == nil || l < 2 {
		return
	}
	end0, start2 := 0, l-1

	for i := 0; i <= start2; {
		t := nums[i]
		switch t {
		case 0:
			swap(end0, i, nums)
			end0++
			i++
		case 1:
			i++
		case 2:
			swap(start2, i, nums)
			start2--
		}
	}
}

func swap(a, b int, nums []int) {
	// fmt.Printf("%p \n", &nums)
	// fmt.Printf("%p \n", &(nums[0]))
	t := nums[a]
	nums[a] = nums[b]
	nums[b] = t
}

// func main() {
// 	arr := []int{2, 0, 2, 1, 1, 0}
// 	// fmt.Printf("%p \n", &arr)
// 	// fmt.Printf("%p \n", &(arr[0]))
// 	// swap(0, 1, arr)
// 	sortColors(arr)
// 	fmt.Println(arr)
// }
