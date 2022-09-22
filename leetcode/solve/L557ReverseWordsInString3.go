package main

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格

//示例 1:
// 输入: "Let's take LeetCode contest"
// 输出: "s'teL ekat edoCteeL tsetnoc"

func reverseWords(s string) string {
	runes := []rune(s)
	from, to := -1, -1
	for index, ru := range runes {
		if ru == ' ' {
			from, to = to, index
			reverseRuneSlice(runes[from+1 : to])
		}
	}
	from, to = to, len(runes)
	reverseRuneSlice(runes[from+1 : to])

	return string(runes)
}

func reverseRuneSlice(ra []rune) {
	for i, j := 0, len(ra)-1; i < j; i, j = i+1, j-1 {
		ra[i], ra[j] = ra[j], ra[i]
	}
}

// type runeArray []rune

// func (ra runeArray) Len() int {
// 	return len(ra)
// }
// func (ra runeArray) Less(i, j int) bool {
// 	return true
// }
// func (ra runeArray) Swap(i, j int) {
// 	ra[i], ra[j] = ra[j], ra[i]
// }

// func main() {
// 	reverseWords("Let's take LeetCode contest")
// }
