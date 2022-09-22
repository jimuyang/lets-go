package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
// 	merge(input)
// }

// merge
// Given a collection of intervals, merge all overlapping intervals.

// Example 1:
// Input: [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)
	// 按第一个元素排序？
	// sort.Sort
	var temp Temp = intervals
	sort.Sort(temp)
	var a int
	fmt.Println("a " + string(a))
	for i := 0; i < len(intervals); i++ {
		if i == 0 {
			result = append(result, intervals[0])
			a = intervals[0][1]
			continue
		}

		if intervals[i][0] <= a {
			// 出现重合 需要合并
			// result[i-1][1] = intervals[i][1]
			if intervals[i][1] < a {
				result[len(result)-1][1] = a
			} else {
				result[len(result)-1][1] = intervals[i][1]
				a = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
			a = intervals[i][1]
		}
	}
	// fmt.Println(temp)
	return result
}

// Temp 对Go的类型还是不懂
type Temp [][]int

func (t Temp) Len() int {
	return len(t)
}

func (t Temp) Less(i, j int) bool {
	return t[i][0] < t[j][0]
}

func (t Temp) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
