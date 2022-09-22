package main

func maxArea(height []int) int {
	// 双指针法 每次移动较矮的指针
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		if height[i] < height[j] {
			area := height[i] * (j - i)
			if area > max {
				max = area
			}
			i++
		} else {
			area := height[j] * (j - i)
			if area > max {
				max = area
			}
			j--
		}
	}
	return max
}
