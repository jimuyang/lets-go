package main

import (
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	if n <= 0 || k <= 0 {
		return ""
	}
	// 计算n的阶乘
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial = factorial * i
	}

	var result strings.Builder

	// 标记数组
	mark := make([]bool, n)

	remain := k
	// 再一层层除下去
	for i := n; factorial >= 1; i-- {
		// 被除 商
		quotient := remain / factorial
		// 余数
		remain = remain % factorial
		if i == n {
			factorial = factorial / i
			continue
		}

		if remain == 0 && quotient == 0 {
			// 直接倒序
			for j := n - 1; j >= 0 && !mark[j]; j-- {
				result.WriteString(strconv.Itoa(j + 1))
			}
			return result.String()
		}

		// 根据商得到mark中第几个还没用过的值
		target := 0
		beforeTarget := 0
		for j := 0; j < n; j++ {
			if !mark[j] {
				quotient--
				if quotient == 0 {
					beforeTarget = j + 1
				}

				if quotient < 0 {
					mark[j] = true
					target = j + 1
					break
				}
			}

		}

		// 余数为0时 使用beforeTarget剩下倒序 并可以返回
		if remain == 0 {
			result.WriteString(strconv.Itoa(beforeTarget))
			mark[beforeTarget-1] = true
			mark[target-1] = false
			for j := n - 1; j >= 0; j-- {
				if !mark[j] {
					result.WriteString(strconv.Itoa(j + 1))
				}
			}
			return result.String()
		} else {
			result.WriteString(strconv.Itoa(target))
		}
		// 最后再除i
		factorial = factorial / i
	}
	return ""
}

// func main() {
// 	fmt.Println(getPermutation(3, 4))
// }
