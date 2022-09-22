package main

func plusOne(digits []int) []int {
	if digits == nil || len(digits) < 1 {
		return digits
	}
	l := len(digits)
	jin := 1
	for i := l - 1; i >= 0; i-- {
		r := digits[i] + jin
		if r >= 10 {
			jin = 1
			r = r - 10
		} else {
			jin = 0
		}
		digits[i] = r
	}
	if jin == 1 {
		// 扩容
		result := make([]int, l+1)
		result[0] = 1
		copy(result[1:], digits)
		return result
	}
	return digits
}
