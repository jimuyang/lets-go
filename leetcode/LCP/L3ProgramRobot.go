package main

// https://leetcode-cn.com/problems/programmable-robot/

// 关键在于无限循环 还有一个条件是 至少一个U一个R
func robot(command string, obstacles [][]int, x int, y int) bool {
	// 更换obstacle表示
	obstacleMap := make(map[int64]bool)
	for _, val := range obstacles {
		obstacleMap[int64(val[0])*1e9+int64(val[1])] = true
	}

	posX, posY := 0, 0
	for posX <= x && posY <= y {
		for _, ch := range command {
			switch ch {
			case 'U':
				posY++
			case 'R':
				posX++
			default:
				panic("unknown command")
			}
			// check
			if obstacleMap[int64(posX)*1e9+int64(posY)] == true {
				return false
			}

			// for _, val := range obstacles {
			// 	if val[0] == posX && val[1] == posY {
			// 		fmt.Println(val)
			// 		return false
			// 	}
			// }
			if posX == x && posY == y {
				return true
			}
		}
	}
	return false
}

// 可以使用周期来优化
