package main

import "container/list"

// 返回 A 的最短的非空连续子数组的长度，该子数组的和至少为 K 。
// 如果没有和至少为 K 的非空子数组，返回 -1 。

func shortestSubarray(A []int, K int) int {
	result := -1

	sums := make([]int, len(A)+1)
	sums[0] = 0
	for i := 1; i < len(A)+1; i++ {
		sums[i] = sums[i-1] + A[i-1]
	}
	var stack list.List
	for i := 0; i < len(A)+1; i++ {
		for stack.Len() > 0 && sums[stack.Front().Value.(int)] >= sums[i] {
			stack.Remove(stack.Front())
		}
		stack.PushFront(i)
		// 计算栈内首尾是否大于K
		for sums[stack.Front().Value.(int)]-sums[stack.Back().Value.(int)] >= K {
			l := stack.Front().Value.(int) - stack.Back().Value.(int)
			if result == -1 || result > l {
				result = l
			}
			// 用过就丢了
			stack.Remove(stack.Back())
		}
	}
	return result
}

// func main() {
// 	shortestSubarray([]int{1}, 1)
// }
