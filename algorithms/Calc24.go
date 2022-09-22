package main

import (
	"container/list"
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
)

// func main() {
// 	// fmt.Println(8.0 / (3.0 - 8.0/3.0))
// 	// fmt.Println(calcExpression(" 8 3 2 6 * + 5 / - 2 2 + +"))
// 	// testPick()
// 	calc24([]int{3, 3, 7, 7})
// }

func calc24(four []int) {
	// 4个数字和3个符号全排列
	charPerms := pickChars([]rune{'*', '+', '-', '/'}, 3)
	for _, charPerm := range charPerms {
		perm := make([]int, 0)
		for _, n := range four {
			perm = append(perm, n+'0')
		}
		for _, pch := range charPerm {
			perm = append(perm, int(pch))
		}
		// fmt.Println(perm)
		// 拿到一组 数字+符号
		allPerms := permuteUnique(perm)
		for _, perm := range allPerms {
			// 转化为exp
			var sb strings.Builder
			for _, ch := range perm {
				sb.WriteRune(rune(ch))
			}
			exp := sb.String()

			val, err := calcExpression(exp)
			if err == nil && math.Abs(val-24.0) < 0.001 {
				fmt.Println(exp, "=", val)
			}
		}
	}
}

// 从所有操作符中挑选k个 每个操作符可重复使用 输出每个char的个数
func pick(chars []rune, k int) [][]int {
	// 实际上是一种分布 k个小球放入len(chars)个槽中
	// 递归来实现 第一个槽中分别放入 0,1,2...k个小球
	result := make([][]int, 0)
	if k == 0 {
		result = append(result, make([]int, len(chars)))
		return result
	}
	if len(chars) == 1 {
		putAll := make([]int, 1)
		putAll[0] = k
		result = append(result, putAll)
		return result
	}
	for i := 0; i <= k; i++ {
		left := pick(chars[1:], k-i)
		for j, v := range left {
			left[j] = append(v, i)
		}
		// 顺序和chars顺序相反
		result = append(result, left...)
	}
	return result
}

func pickChars(chars []rune, k int) [][]rune {
	result := make([][]rune, 0)
	charNum := pick(chars, k)
	for _, res := range charNum {
		// var sb strings.Builder
		one := make([]rune, 0)
		for i, num := range res {
			for j := 0; j < num; j++ {
				// sb.WriteRune(chars[len(chars)-i-1])
				one = append(one, chars[len(chars)-i-1])
			}
		}
		result = append(result, one)
	}
	return result
}

func testPick() {
	chars := []rune{'*', '+', '-', '/'}
	k := 3
	result := pick(chars, k)
	fmt.Println(len(result))
	fmt.Println(pickChars(chars, k))
}

// 计算后缀表达式的值 认为0是10; 8 3 2 6 * + 5 / - 2 2 ^ +
func calcExpression(exp string) (float64, error) {
	var stack list.List
	for _, ch := range []rune(exp) {
		if ch >= '0' && ch <= '9' {
			if ch == '0' {
				stack.PushFront(float64(10))
			} else {
				stack.PushFront(float64(ch - '0'))
			}
			continue
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			if stack.Len() == 0 {
				return 0.0, errors.New("invalid expression")
			}
			op1 := stack.Front().Value.(float64)
			stack.Remove(stack.Front())
			if stack.Len() == 0 {
				return 0.0, errors.New("invalid expression")
			}
			op2 := stack.Front().Value.(float64)
			stack.Remove(stack.Front())
			var result float64
			switch ch {
			case '+':
				result = op2 + op1
			case '-':
				result = op2 - op1
			case '*':
				result = op2 * op1
			case '/':
				result = op2 / op1
			}
			stack.PushFront(result)
		}
	}
	if stack.Len() != 1 {
		return 0.0, errors.New("invalid expression")
	}
	return stack.Front().Value.(float64), nil
}

// 不重复的全排列
func permuteUnique(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	result := make([][]int, 0)
	result = append(result, append(make([]int, 0), nums...))

	for {
		// 从尾部开始寻找降序结构
		i := len(nums) - 1
		for ; i >= 1 && nums[i] <= nums[i-1]; i-- {
		}
		if i == 0 {
			// 全降序
			return result
		}
		ti := i - 1
		// 找到第一个比它大的数
		for i < len(nums) && nums[i] > nums[ti] {
			i++
		}
		// 交换
		nums[i-1], nums[ti] = nums[ti], nums[i-1]
		if ti != len(nums)-1 {
			reverseArray(nums[ti+1:])
		}
		result = append(result, append(make([]int, 0), nums...))
	}
}

// 反序
func reverseArray(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
