package main

import "fmt"

// k次交易
func MaxProfitCommon(prices []int, k int) int {
	// 求出所有上升区段 {{price低点, price高点, benefit}}
	allUp := make([][]int, 0)

	start := -1
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			continue
		}
		if prices[i] >= prices[i-1] {
			// 上升
			if start == -1 {
				start = i - 1
			}
			if i == len(prices)-1 {
				allUp = append(allUp, []int{prices[start], prices[i], prices[i] - prices[start]})
			}
		} else {
			// 下降
			if start != -1 {
				allUp = append(allUp, []int{prices[start], prices[i-1], prices[i-1] - prices[start]})
			}
			start = -1
		}
	}
	// fmt.Println(allUp)
	if len(allUp) == 0 || k == 0 {
		return 0
	}

	if k < len(allUp) {
		for i := len(allUp); i != k; i-- {
			// 舍弃或者合并一笔交易 最小的损耗
			minLoss, group := -1, -1
			for i := 0; i < len(allUp)-1; i++ {
				_, loss := mergeOrDiscard(allUp[i], allUp[i+1])
				if minLoss == -1 || loss < minLoss {
					minLoss = loss
					group = i
				}
			}
			// fmt.Println(minLoss, group)
			// 把合并结果替换group
			merge, _ := mergeOrDiscard(allUp[group], allUp[group+1])
			allUp[group] = merge
			allUp = append(allUp[:group+1], allUp[group+2:]...)
			fmt.Println(allUp)
		}
	}

	allBenefit := 0
	for _, se := range allUp {
		allBenefit += se[2]
	}
	return allBenefit
}

// 计算能否合并2笔交易 合并的话损失为多少
func mergeOrDiscard(trade1, trade2 []int) ([]int, int) {
	if trade2[1] >= trade1[1] && trade1[0] <= trade2[0] {
		merge := []int{trade1[0], trade2[1], trade2[1] - trade1[0]}
		loss := trade1[2] + trade2[2] - merge[2]
		return merge, loss
	} else {
		if trade1[2] <= trade2[2] {
			return trade2, trade1[2]
		} else {
			return trade1, trade2[2]
		}
	}
}
