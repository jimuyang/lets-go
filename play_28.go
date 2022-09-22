package main

import "fmt"

var all = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 先不考虑花色
func bigger(zhuang []int, xian []int) bool {
	isDui := func(cards []int) bool {
		return cards[0] == cards[1]
	}
	is28 := func(cards []int) bool {
		return (cards[0] == 2 && cards[1] == 8) || (cards[1] == 2 && cards[0] == 8)
	}

	if is28(zhuang) {
		return true
	} else if is28(xian) {
		return false
	}

	if isDui(zhuang) {
		return true
	} else if isDui(xian) {
		return false
	}

	zhuangDian := zhuang[0] + zhuang[1]
	if zhuangDian >= 10 {
		zhuangDian -= 10
	}

	xianDian := xian[0] + xian[1]
	if xianDian >= 10 {
		xianDian -= 10
	}
	return zhuangDian >= xianDian
}

func play28() {

	total := 0
	win := 0

	for _, z1 := range all {
		for _, z2 := range all {
			for _, x1 := range all {
				for _, x2 := range all {
					zhuang := []int{z1, z2}
					xian := []int{x1, x2}
					total += 1
					if bigger(zhuang, xian) {
						win += 1
					}
				}
			}
		}
	}

	fmt.Println(total, win)
}
