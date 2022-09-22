package main

import "fmt"

// 在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
// 如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/gas-station
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func canCompleteCircuit(gas []int, cost []int) int {
	// 遍历计算经过每个加油站的净值
	totalGas, totalCost := 0, 0
	// 净值保存
	net := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		net[i] = gas[i] - cost[i]
	}
	if totalCost > totalGas {
		return -1
	}
	for start := 0; start < len(gas); {
		// 从start出发
		tank := net[start]
		if tank < 0 {
			start++
			continue
		}
		// 试试能否经过所有站点
		i := start + 1
		for loopBack := false; ; i++ {
			if i >= len(gas) {
				i -= len(gas)
				loopBack = true
			}
			tank += net[i]
			if tank < 0 {
				// 无法到达下一个站点
				if loopBack {
					fmt.Println("loopback")
					return -1
				} else {
					start = i + 1
					break
				}
			}
			if i == start {
				return start
			}
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(canCompleteCircuit(
// 		[]int{3, 1, 1},
// 		[]int{1, 2, 2}))
// }
