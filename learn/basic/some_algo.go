package main

// Add 加法
func Add(a, b int) int {
	return a + b
}

// NonRepeatSubStringMaxLen 字符串的最长不重复字串
func NonRepeatSubStringMaxLen(str string) int {
	lastOccur := make(map[rune]int)
	start := 0
	maxLen := 0
	for i, ch := range []rune(str) {
		if last, ok := lastOccur[ch]; ok && last >= start {
			// 遇到和之前的重复
			start = last + 1
		}
		// 每次计算长度
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccur[ch] = i
	}
	return maxLen
}

// 使用数组要注意避开0
var lastOccur = make([]int, 0xffff)

// NonRepeatSubStringMaxLen1 字符串的最长不重复字串 使用数组优化
func NonRepeatSubStringMaxLen1(str string) int {
	// init
	for i := range lastOccur {
		lastOccur[i] = 0
	}
	start := 0
	maxLen := 0
	for i, ch := range []rune(str) {
		if last := lastOccur[ch] - 1; last >= 0 && last >= start {
			// 遇到和之前的重复
			start = last + 1
		}
		// 每次计算长度
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccur[ch] = i + 1
	}
	return maxLen
}
