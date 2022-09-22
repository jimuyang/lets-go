package main

import (
	"container/list"
)

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水

// 单调栈的做法关键是形成凹结构
func trap(height []int) int {
	var stack list.List
	result := 0

	for i := 0; i < len(height); i++ {
		h := height[i]
		for stack.Len() > 0 && h > height[stack.Front().Value.(int)] {
			// 出栈
			out := stack.Front()
			stack.Remove(out)
			if stack.Len() > 0 {
				// 形成凹
				water := 0
				li := stack.Front().Value.(int)
				lh := height[li]
				di := height[out.Value.(int)]
				if lh < h {
					water = (lh - di) * (i - li - 1)
				} else {
					water = (h - di) * (i - li - 1)
				}
				result += water
			}
		}
		// 入栈
		stack.PushFront(i)
	}
	return result
}

// func main() {
// 	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
// }
