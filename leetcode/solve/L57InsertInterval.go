package main

import (
	"fmt"
)

// Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
// You may assume that the intervals were initially sorted according to their start times.

// Example 1:
// Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
// Output: [[1,5],[6,9]]

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	if intervals == nil || len(intervals) < 1 {
		result = append(result, newInterval)
		return result
	}
	if newInterval == nil {
		return intervals
	}

	// 因为原数组是non-overlapping的 因此需要一个flag来标识需要处理overlap时
	// needCheck := false
	for i := 0; i < len(intervals) || newInterval != nil; i++ {
		// 抉择：当前项还是newInterval
		var target []int
		if newInterval != nil && (i >= len(intervals) || newInterval[0] < intervals[i][0]) {
			target = newInterval
			newInterval = nil
			i--
		} else {
			target = intervals[i]
		}
		// 第一次直接append
		if len(result) == 0 {
			result = append(result, target)
			continue
		}
		// 需要检查是否与前一项重合
		tail := result[len(result)-1][1]
		if target[0] <= tail {
			// 出现重合 需要合并
			if target[1] > tail {
				result[len(result)-1][1] = target[1]
			}
		} else {
			// 没出现重合
			result = append(result, target)
		}
	}
	return result
}

func mergeInterval(first []int, second []int) (canMerge bool, mergeResult []int, after1 []int, after2 []int) {
	if first[1] >= second[0] || second[1] >= first[0] {
		var left, right int
		if first[0] <= second[0] {
			left = first[0]
		} else {
			left = second[0]
		}
		fmt.Println(left)
		if first[1] >= first[1] {
			right = first[1]
		} else {
			right = second[1]
		}
		return true, []int{left, right}, nil, nil
	}
	// 不可以合并的话 输出排序后的结果
	if first[0] < second[0] {
		return false, nil, first, second
	} else {
		return false, nil, second, first
	}
}

// func main() {
// 	input := [][]int{{1, 3}, {6, 9}}
// 	fmt.Println(insert(input, []int{2, 5}))
// }
