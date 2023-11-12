package main

/*
You are given an array routes representing bus routes where routes[i] is a bus route that the ith bus repeats forever.

For example, if routes[0] = [1, 5, 7], this means that the 0th bus travels in the sequence 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... forever.
You will start at the bus stop source (You are not on any bus initially), and you want to go to the bus stop target. You can travel between bus stops by buses only.

Return the least number of buses you must take to travel from source to target. Return -1 if it is not possible.



Example 1:

Input: routes = [[1,2,7],[3,6,7]], source = 1, target = 6
Output: 2
Explanation: The best strategy is take the first bus to the bus stop 7, then take the second bus to the bus stop 6.
Example 2:

Input: routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
Output: -1

*/

/*

 */

func numBusesToDestination(routes [][]int, source int, target int) int {

	// 换个思路 看bus的可达性

	stopBus := make(map[int][]int)
	for busI, bus := range routes {
		for _, stop := range bus {
			stopBus[stop] = append(stopBus[stop], busI)
		}
	}
	sourceBus := stopBus[source]
	targetBus := stopBus[target]

	targetBusHit := make(map[int]bool)
	for _, bus := range targetBus {
		targetBusHit[bus] = true
	}

	// 通过stopBus可以知道bus换乘关系
	bus2Bus := make(map[int][]int)
	for _, busList := range stopBus {
		for _, busI := range busList {
			for _, toBus := range busList {
				if toBus != busI {
					bus2Bus[busI] = append(bus2Bus[busI], toBus)
				}
			}
		}
	}

	canBus := make(map[int]int)
	for _, bus := range sourceBus {
		canBus[bus] = 1
	}
	newBus := sourceBus

	for len(newBus) > 0 {
		temp := make([]int, 0)

		for _, bus := range newBus {
			if targetBusHit[bus] {
				return canBus[bus]
			}
			i := canBus[bus]
			for _, nextBus := range bus2Bus[bus] {
				if canBus[nextBus] == 0 {
					canBus[nextBus] = i + 1
					temp = append(temp, nextBus)
				}

			}

		}
		newBus = temp
	}
	return -1

	// 可达分析   时间超了
	// 核心问题是按照啥顺序做可达分析？

	//// 每个站能到哪些站
	//stop2Stop := make(map[int][]int)
	//for _, bus := range routes {
	//	for _, stop := range bus {
	//		for _, toStop := range bus {
	//			if toStop != stop {
	//				stop2Stop[stop] = append(stop2Stop[stop], toStop)
	//			}
	//		}
	//		stop2Stop[stop] = distinctInt(stop2Stop[stop])
	//
	//	}
	//}
	//
	//// 初始值
	//canArrive := make(map[int]int)
	//canArrive[source] = 1
	//newStops := []int{source}
	//for len(newStops) > 0 {
	//	temp := make([]int, 0)
	//
	//	for _, stop := range newStops {
	//		i := canArrive[stop] // 从这个站出发
	//		for _, nextStop := range stop2Stop[stop] {
	//			if canArrive[nextStop] == 0 {
	//				canArrive[nextStop] = i + 1
	//				temp = append(temp, nextStop)
	//			}
	//			if nextStop == target {
	//				return canArrive[target] - 1
	//			}
	//
	//		}
	//	}
	//	newStops = temp
	//}
	//return canArrive[target] - 1
}

func mergeBus(routes [][]int) []int {
	res := make([]int, 0)
	for _, bus := range routes {
		res = append(res, bus...)
	}
	return distinctInt(res)
}

func distinctInt(arr []int) []int {
	hit := make(map[int]bool, len(arr))
	for _, i := range arr {
		hit[i] = true
	}
	res := make([]int, 0, len(hit))
	for i := range hit {
		res = append(res, i)
	}
	return res
}
