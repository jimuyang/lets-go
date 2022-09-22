package main

import (
	"strconv"
	"strings"
)

// 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
// 示例:
// 输入: "25525511135"
// 输出: ["255.255.11.135", "255.255.111.35"]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/restore-ip-addresses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func restoreIPAddresses(s string) []string {
	// 其实就是找3个分割点 使得分割成4个子串 每个子串 >= 0 且 <= 255
	if len(s) < 4 {
		return []string{}
	}
	l := len(s)
	result := make([]string, 0)
	// 25525511135 为例
	// 第一个分割点很简单 对应坐标012位置 要求后续至少3位
	n1, n2, n3 := 0, 0, 0
	for f1, e1 := 0, 3; f1 < e1 && l-f1 >= 4; f1++ {
		n1 = n1*10 + int(s[f1]) - int('0')
		if n1 > 255 {
			break
		}
		if n1 == 0 {
			e1 = f1
		}
		// 第二个分割点
		n2 = 0
		for f2, e2 := f1+1, f1+4; f2 < e2 && l-f2 >= 3; f2++ {
			n2 = n2*10 + int(s[f2]) - int('0')
			if n2 > 255 {
				break
			}
			if n2 == 0 {
				e2 = f2
			}
			// 第三个分割点
			n3 = 0
			for f3, e3 := f2+1, f2+4; f3 < e3 && l-f3 >= 2; f3++ {
				n3 = n3*10 + int(s[f3]) - int('0')
				if n3 > 255 {
					break
				}
				if n3 == 0 {
					e3 = f3
				}
				// 剩下超过3位
				if l-f3 > 4 {
					continue
				}
				// 剩下开头为0且超过1位
				if s[f3+1] == '0' && l-f3 > 2 {
					continue
				}
				// 剩下超过255
				left, _ := strconv.Atoi(s[f3+1:])
				if left > 255 {
					continue
				}
				var builder strings.Builder
				builder.WriteString(s[:f1+1])
				builder.WriteRune('.')
				builder.WriteString(s[f1+1 : f2+1])
				builder.WriteRune('.')
				builder.WriteString(s[f2+1 : f3+1])
				builder.WriteRune('.')
				builder.WriteString(s[f3+1:])
				result = append(result, builder.String())
			}
		}
	}
	return result
}

// func main() {
// 	fmt.Println(restoreIPAddresses("25525511135"))
// }
