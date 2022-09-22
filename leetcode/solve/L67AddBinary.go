package main

import (
	"container/list"
	"strconv"
)

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	la, lb := len(a), len(b)
	// 更小的长度 更长的string
	minL, maxL, longer, shorter := la, lb, b, a
	if la > lb {
		minL, maxL, longer, shorter = lb, la, a, b
	}

	result := list.New()
	var jin byte
	for i := maxL - 1; i >= 0; i-- {
		si := i - (maxL - minL)
		var r byte
		if si >= 0 {
			r = shorter[si] - '0' + longer[i] - '0' + jin
		} else {
			r = longer[i] - '0' + jin
		}

		if r >= 2 {
			r = r - 2
			jin = 1
		} else {
			jin = 0
		}
		result.PushFront(r)
	}
	if jin > 0 {
		result.PushFront(jin)
	}

	// var builder strings.Builder
	str := ""
	for e := result.Front(); e != nil; e = e.Next() {
		str += strconv.Itoa(int(e.Value.(byte)))
		// builder.WriteByte(e.Value.(byte))
	}
	return str
}

// func main() {
// 	fmt.Println(addBinary("11", "1"))
// }
