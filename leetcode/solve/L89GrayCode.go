package main

//格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。
// 给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。格雷编码序列必须以 0 开头。
// 输入: 2
// 输出: [0,1,3,2]
// 解释:
// 00 - 0
// 01 - 1
// 11 - 3
// 10 - 2

// 对于给定的 n，其格雷编码序列并不唯一。
// 例如，[0,2,3,1] 也是一个有效的格雷编码序列。

// 00 - 0
// 10 - 2
// 11 - 3
// 01 - 1

// 递推试试吧
func grayCode(n int) []int {
	if n <= 0 {
		return []int{0}
	}
	result := make([]int, 1<<uint(n))
	// fmt.Println(len(result))
	l := 0
	for i := 0; i <= n; i++ {
		if i == 0 {
			result[0] = 0
			l = 1
			continue
		}
		for j := l - 1; j >= 0; j-- {
			result[2*l-1-j] = 1<<uint(i-1) | result[j]
		}
		l = l * 2
	}
	return result
}

// func main() {
// 	fmt.Println(grayCode(2))
// }
