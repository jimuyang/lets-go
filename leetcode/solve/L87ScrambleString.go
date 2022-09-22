package main

//https://leetcode-cn.com/problems/scramble-string/
// 看了别人的解法 得到的重要经验：
// 递归可以使用动归优化 重要的是如果保存结果

// 递归试试
func isScramble(s1 string, s2 string) bool {
	l := len(s1)
	if l != len(s2) {
		return false
	}
	if l <= 1 {
		return s1 == s2
	}
	// fail-fast 统计字符出现次数
	chMap1 := make(map[byte]int)
	chMap2 := make(map[byte]int)
	for i := 0; i < l; i++ {
		chMap1[s1[i]]++
		chMap2[s2[i]]++
	}
	for k, v := range chMap1 {
		if v != chMap2[k] {
			return false
		}
	}

	// 遍历分割点
	for i := 1; i < l; i++ {
		if (isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:])) ||
			(isScramble(s1[:i], s2[l-i:]) && isScramble(s1[i:], s2[:l-i])) {
			return true
		}
	}
	return false
}

// func main() {
// 	fmt.Println(isScramble("great", "rgate"))
// }
