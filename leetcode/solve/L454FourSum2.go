package main

// 给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
// 为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。

// 目测不用去重
func fourSumCount(A []int, B []int, C []int, D []int) int {
	result := 0
	// D转为 map[value]indexes
	Dmap := make(map[int][]int)
	for i, v := range D {
		if indexes, ok := Dmap[v]; ok {
			Dmap[v] = append(indexes, i)
		} else {
			Dmap[v] = []int{i}
		}
	}
	for _, av := range A {
		for _, bv := range B {
			for _, cv := range C {
				need := 0 - av - bv - cv
				if dis, ok := Dmap[need]; ok {
					result += len(dis)
				}
			}
		}
	}
	return result
}

// 上面复杂度为n^3 试试扩大map

func fourSumCount1(A []int, B []int, C []int, D []int) int {
	result := 0
	CDMap := make(map[int]int)
	// CD转为 map[c+d]count
	for _, cv := range C {
		for _, dv := range D {
			CDMap[cv+dv]++
		}
	}
	for _, av := range A {
		for _, bv := range B {
			need := 0 - av - bv
			result += CDMap[need]
		}
	}
	return result
}
