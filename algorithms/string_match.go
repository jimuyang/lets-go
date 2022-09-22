package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 字符串匹配的几种算法

// 暴力搜索
func forceSearch(source, pattern string) int {
	src, pat := []rune(source), []rune(pattern)
	ls, lp := len(src), len(pat)
	for i := 0; i <= ls-lp; i++ {
		// 首字母是否相等
		if src[i] == pat[0] {
			match := true
			// 继续比较剩下的字母是否相等
			for j := 1; j < lp; j++ {
				if pat[j] != src[i+j] {
					match = false
					break
				}
			}
			if match {
				return i
			}
		}
	}
	return -1
}

// 基于有限状态机
func smSearch(source, pattern string) int {
	src, pat := []rune(source), []rune(pattern)
	ls, lp := len(src), len(pat)
	status := 0
	sm := buildStatusMachine(pat)
	// for i, ch := range src {
	for i := 0; i < ls; i++ {
		// 如果ch不在状态机的分支中 此时status=0
		status = sm[status][src[i]]
		if status == lp {
			return i - lp + 1
		}
	}
	return -1
}

// build有限状态机
func buildStatusMachine(pat []rune) []map[rune]int {
	lp := len(pat)
	// 状态机
	sm := make([]map[rune]int, lp+1)
	for i := 0; i < lp+1; i++ {
		mp := make(map[rune]int)
		sm[i] = mp
	}
	// 解决和开头重复的问题 也就是KMP的最大长度表
	repeat := buildRepeat(pat)
	// 基于repeat数组可以给到状态机的回头跳转
	for i, rep := range repeat {
		sm[i+1][pat[rep]] = rep + 1
	}
	// 按pat顺序写入状态机 A-B-A-C 同时覆盖掉部分回头跳转
	for i := 0; i < lp; i++ {
		sm[i][pat[i]] = i + 1
	}
	return sm
}

// KMP 其实就是在暴力搜索中 加速了外层循环
func kmpSearch(source, pattern string) int {
	src, pat := []rune(source), []rune(pattern)
	ls, lp := len(src), len(pat)
	repeat := buildRepeat(pat)
	for i := 0; i <= ls-lp; i++ {
		// 首字母是否相等
		if src[i] == pat[0] {
			match := true
			// 继续比较剩下的字母是否相等
			for j := 1; j < lp; j++ {
				if pat[j] != src[i+j] {
					match = false
					// 这里加速外层循环
					i += j - repeat[j-1] - 1
					break
				}
			}
			if match {
				return i
			}
		}
	}
	return -1
}

// 广泛使用的尾部匹配高效算法
func bmSearch(source, pattern string) int {
	src, pat := []rune(source), []rune(pattern)
	ls, lp := len(src), len(pat)

	// 每个字符在pat中最后出现的位置
	// lastPos := make(map[rune]int)
	// for i, ch := range pat {
	// 	lastPos[ch] = i
	// }
	lastPosArr := make([]int, 256)
	for i := 0; i < 256; i++ {
		lastPosArr[i] = -1
	}
	for i, ch := range pat {
		lastPosArr[ch] = i
	}

	// sameSuffix[i]: pat[:i+1]和pat的公共后缀长度
	sameSuffix := make([]int, lp)
	for i := lp; i > 0; i-- {
		sameSuffix[i-1] = sameSuffixLen(pat, pat[:i])
	}
	// goodSuffix[i]: 当goodSuffix长度为i时 应当右移步数
	goodSuffix := make([]int, lp+1)
	for i := 0; i < lp; i++ {
		goodSuffix[i] = lp
	}
	// 正序遍历sameSuffix 目的是减少移动步数 防止漏过
	for i, ssl := range sameSuffix {
		goodSuffix[ssl] = lp - i - 1
	}
	// 开始
	for i := lp - 1; i < ls; {
		// 从尾部开始比较
		for j := i; j > i-lp; j-- {
			patI := lp - (i - j) - 1
			if src[j] != pat[patI] {
				// badChar出现
				pos := 0
				// if v, ok := lastPos[src[j]]; ok {
				if v := lastPosArr[src[j]]; v >= 0 {
					pos = v
				} else {
					pos = -1
				}
				// badChar对齐需要移动
				move := patI - pos
				if j < i {
					// goodSuffix出现
					if move < goodSuffix[i-j] {
						move = goodSuffix[i-j]
					}
				}
				i += move
				break
			}
			if j == i-lp+1 {
				return j
			}
		}
	}
	return -1
}

// 公共后缀长度
func sameSuffixLen(pat, patPart []rune) int {
	l, lp := len(pat), len(patPart)
	for i := 0; i < lp; i++ {
		if pat[l-1-i] != patPart[lp-1-i] {
			return i
		}
	}
	return lp
}

// 其实是最大长度表
func buildRepeat(pat []rune) []int {
	repeat := make([]int, len(pat))
	for i := 1; i < len(pat); i++ {
		if pat[i] == pat[repeat[i-1]] {
			repeat[i] = repeat[i-1] + 1
		}
	}
	return repeat
}

func main2() {
	rand.Seed(time.Now().UnixNano())

	source := genString(1000)
	fmt.Println(source)
	var ft, st, kt, bt int64
	var t1 time.Time
	for i := 0; i < 100; i++ {
		pattern := genString(100)
		fmt.Println(pattern)

		t1 = time.Now()
		fmt.Println(smSearch(source, pattern))
		st = st + time.Since(t1).Nanoseconds()

		t1 = time.Now()
		fmt.Println(kmpSearch(source, pattern))
		kt = kt + time.Since(t1).Nanoseconds()

		t1 = time.Now()
		fmt.Println(bmSearch(source, pattern))
		bt = bt + time.Since(t1).Nanoseconds()

		t1 = time.Now()
		fmt.Println(forceSearch(source, pattern))
		ft = ft + time.Since(t1).Nanoseconds()
	}
	fmt.Printf("暴力搜索：%v\n状态机  ：%v\nKMP     ：%v\nBM      ：%v\n", ft, st, kt, bt)
}

func genString(n int) string {
	chars := make([]rune, 0)
	for i := 0; i < 5; i++ {
		chars = append(chars, 'a'+rune(i))
	}

	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteRune(chars[rand.Intn(len(chars))])
	}
	return builder.String()
}
