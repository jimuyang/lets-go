package main

import (
	"container/list"
)

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。

// 有点像 降水那题
func largestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}
	result := 0
	stack := list.New()
	for _, height := range heights {
		if stack.Len() == 0 {
			stack.PushBack(&HeightCount{height, 1})
			continue
		}
		max := 0
		// 出栈直到栈顶 <= height
		for top := stack.Back(); top.Value.(*HeightCount).height > height; {
			hc := top.Value.(*HeightCount)
			if hc.count > max {
				max = hc.count
			}
			// 计算面积
			area := hc.count * hc.height
			if area > result {
				result = area
			}
			stack.Remove(top)
			if stack.Len() == 0 {
				break
			}
			top = stack.Back()
		}

		// 栈内所有元素计数加1
		if stack.Len() > 0 {
			for node := stack.Front(); node != nil; node = node.Next() {
				hc := node.Value.(*HeightCount)
				hc.count++
			}
		}
		// 入栈
		stack.PushBack(&HeightCount{height, max + 1})
	}
	for top := stack.Back(); ; {
		// 计算面积
		hc := top.Value.(*HeightCount)
		area := hc.count * hc.height
		if area > result {
			result = area
		}
		stack.Remove(top)
		if stack.Len() == 0 {
			break
		}
		top = stack.Back()
	}
	return result
}

// HeightCount 高度统计
type HeightCount struct {
	height int
	count  int
}

// func main() {
// 	fmt.Println(largestRectangleArea2([]int{4, 2, 0, 3, 2, 4, 3, 4}))
// }

// 单调栈试一下
func largestRectangleArea2(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}
	result := 0
	stack := list.New()
	stack.PushBack(-1)
	for i := 0; i < len(heights); i++ {
		top := stack.Back().Value.(int)
		for top != -1 && heights[i] <= heights[top] {
			stack.Remove(stack.Back())
			newTop := stack.Back().Value.(int)
			area := (i - newTop - 1) * heights[top]
			if area > result {
				result = area
			}
			top = newTop
		}
		stack.PushBack(i)
	}
	// 清空单调栈
	for stack.Len() > 1 {
		top := stack.Back().Value.(int)
		stack.Remove(stack.Back())
		newTop := stack.Back().Value.(int)
		area := (len(heights) - newTop - 1) * heights[top]
		if area > result {
			result = area
		}
	}
	return result
}

// // 正序 逆序遍历各一次
// func largestRectangleArea2(heights []int) int {
// 	left := make([]int, len(heights))
// 	right := make([]int, len(heights))

// 	last := 0
// 	// 正序遍历
// 	for i := 0; i < len(heights); i++ {
// 		if i != 0 && heights[i] <= last {
// 			left[i] = left[i-1] + 1
// 		}
// 		last = heights[i]
// 	}

// 	// 倒序遍历
// 	last = 0
// 	for i := len(heights) - 1; i >= 0; i-- {
// 		if i != len(heights)-1 && heights[i] <= last {
// 			right[i] = right[i+1] + 1
// 		}
// 		last = heights[i]
// 	}
// 	result := 0
// 	for i := 0; i < len(heights); i++ {
// 		r := heights[i] * (left[i] + right[i] + 1)
// 		if r > result {
// 			result = r
// 		}
// 	}
// 	return result
// }
