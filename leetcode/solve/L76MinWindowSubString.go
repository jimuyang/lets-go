package main

import (
	"container/list"
)

// 给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。

// 示例：

// 输入: S = "ADOBECODEBANC", T = "ABC"
// 输出: "BANC"
// 说明：

// 如果 S 中不存这样的子串，则返回空字符串 ""。
// 如果 S 中存在这样的子串，我们保证它是唯一的答案。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-window-substring

func wBack(window *list.List) position {
	return window.Back().Value.(position)
}
func wFront(window *list.List) position {
	return window.Front().Value.(position)
}

// 提示的是滑动窗口的做法
func minWindow(s string, t string) string {
	result := ""

	charMap := make(map[rune]int)
	for _, ch := range t {
		// 对不存在的key go中的map返回0值
		charMap[ch] = charMap[ch] + 1
	}
	uniTotal := len(charMap)
	// fmt.Println("unicode in t:", uniTotal)

	// onlyOne := false
	// if uniTotal == 0 {
	// 	return ""
	// } else if uniTotal == 1 {
	// 	onlyOne = true
	// }

	// 遍历看看
	// uniCount := 0        // 滑动窗口内不同字符数
	var window = list.New() // 滑动窗口
	countMap := make(map[rune]int)
	validCount := 0

	for index, ch := range s {
		_, exist := charMap[ch]
		if !exist {
			continue
		}
		// if onlyOne {
		// 	var sb strings.Builder
		// 	sb.WriteRune(ch)
		// 	return sb.String()
		// }
		countMap[ch]++
		window.PushBack(position{ch, index})

		if countMap[ch] == charMap[ch] {
			// 认为一个字母达到了要求
			validCount++
		}

		if validCount == uniTotal {
			// 所有字母数量都够了
			// fmt.Println(wFront(window).index, wBack(window).index+1)
			temp := s[wFront(window).index : wBack(window).index+1]
			if result == "" || len(result) > len(temp) {
				result = temp
			}
			// 窗口内有所有的字符 此时从前部开始缩小窗口 直到缺失为止
			for e := window.Front(); e != nil; {
				pos := e.Value.(position)
				window.Remove(e)
				countMap[pos.ch]--
				if countMap[pos.ch] == charMap[pos.ch]-1 {
					// 缺失ch
					validCount--
					// 此时长度为
					if window.Len() == 0 {
						temp = s[pos.index : pos.index+1]
					} else {
						temp = s[pos.index : wBack(window).index+1]
					}
					if result == "" || len(result) > len(temp) {
						result = temp
					}
					break
				}
				e = window.Front()
			}
		}
	}

	return result
}

type position struct {
	ch    rune
	index int
}

// func main() {
// 	// fmt.Println(reflect.TypeOf('你'))
// 	// fmt.Println(minWindow("ADOBECODEBANC", "ABCA"))
// 	fmt.Println(minWindow("wubcadbcwerwtrwrwerqa", "aba"))
// }
