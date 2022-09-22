package main

import "fmt"

// 高楼扔鸡蛋
// 一栋100层的大楼，给你一模一样的2个鸡蛋。已知存在一个临界楼层会让鸡蛋摔破，低于这个楼层鸡蛋不会破，高于这个楼层鸡蛋一定破，鸡蛋不摔破时可以捡回来重新扔。
// 请你提供一个扔鸡蛋的策略来找到这个临界楼层，而且这个扔法在最坏情况下有最少的扔鸡蛋次数

// 如果推广到M层高 N个鸡蛋呢？

func throw2EggsFrom100() int {
	// dp[i][j] 代表j+1层大楼 i+1个鸡蛋的最少次数
	dp := make([][]int, 2)
	dp[0] = make([]int, 101)
	dp[1] = make([]int, 101)

	// 1个鸡蛋
	for i := 0; i <= 100; i++ {
		dp[0][i] = i
	}

	// 2个鸡蛋
	dp[1][0] = 0 // 0层楼扔0次
	dp[1][1] = 1 // 1层楼扔1次
	for i := 2; i <= 100; i++ {
		min := dp[1][i-1] + 1
		minFirst := 0
		// 第1个鸡蛋分别从1，2，...i层楼丢下去
		for j := 1; j <= i; j++ {
			// 假如碎了
			t1 := dp[0][j-1]
			// 假如没碎
			t2 := dp[1][i-j]
			// 取最坏情况
			t := t1 + 1
			if t1 < t2 {
				t = t2 + 1
			}
			// 最好扔法
			if min >= t {
				min = t
				minFirst = j
			}
		}
		fmt.Printf("%d层楼 第一个鸡蛋从%d层扔下去 最坏需要%d次 \n", i, minFirst, min)
		dp[1][i] = min
	}
	return dp[1][100]
}

// func main() {
// 	fmt.Println(throw2EggsFrom100())
// 	fmt.Println(commonThrowEgg(100, 3))
// }

func commonThrowEgg(height int, eggCount int) int {
	if height < 1 {
		return 0
	}
	if eggCount <= 1 {
		return height
	}
	dp := make([][]int, eggCount)
	for i := 0; i < eggCount; i++ {
		dp[i] = make([]int, height+1)
	}
	// 初始化部分数据
	// 1个鸡蛋
	for i := 0; i <= height; i++ {
		dp[0][i] = i
	}
	// 1层楼 0层楼
	for i := 0; i < eggCount; i++ {
		dp[i][0] = 0
		dp[i][1] = 1
	}
	// 2个鸡蛋起
	for i := 1; i < eggCount; i++ {
		// 2层楼起
		for j := 2; j <= height; j++ {
			min := dp[i][j-1] + 1
			// 第1个鸡蛋分别从1，2，...j层楼丢下去
			for k := 1; k <= j; k++ {
				// 假如碎了
				t1 := dp[i-1][k-1]
				// 假如没碎
				t2 := dp[i][j-k]
				// 取最坏情况
				t := t1 + 1
				if t1 < t2 {
					t = t2 + 1
				}
				// 最好扔法
				if min >= t {
					min = t
				}
			}
			dp[i][j] = min
		}
	}
	return dp[eggCount-1][height]
}
