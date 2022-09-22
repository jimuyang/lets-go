package main

import "unicode"

// 给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
// 如果不存在最后一个单词，请返回 0 。
// 说明：一个单词是指由字母组成，但不包含任何空格的字符串。
// 示例:
// 输入: "Hello World"
// 输出: 5

func lengthOfLastWord(s string) int {
	// // 用string.Field试试
	// fields := strings.Fields(s)
	// if len(fields) > 0 {
	// 	return len(fields[len(fields)-1])
	// }
	// return 0

	// 逆向遍历呢？
	start := false
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		// fmt.Println(reflect.TypeOf(t))
		if !start && !unicode.IsSpace(rune(s[i])) {
			// start!
			start = true
		}
		if start {
			if unicode.IsSpace(rune(s[i])) {
				break
			} else {
				length++
			}
		}
	}
	return length
}

// func main() {
// 	lengthOfLastWord("aaa你好a")
// }
