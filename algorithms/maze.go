package main

import (
	"container/list"
	"fmt"
	"strings"
)

type coord struct {
	x, y int
}

// 深度优先算法
func breadFirst(maze [][]int, start coord, dest coord) {
	X, Y := len(maze), len(maze[0])
	mark := make([][]int, X)
	for i := 0; i < X; i++ {
		mark[i] = make([]int, Y)
	}
	var queue list.List
	queue.PushBack(start)
	mark[start.x][start.y] = 1

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		check := front.Value.(coord)
		if check.x == dest.x && check.y == dest.y {
			fmt.Println(mark)
			var sb strings.Builder
			for current, step := dest, mark[dest.x][dest.y]; step > 0; {
				sb.WriteString(fmt.Sprintf("(%d, %d) <- ", current.x, current.y))
				around := []coord{
					{current.x - 1, current.y},
					{current.x, current.y - 1},
					{current.x + 1, current.y},
					{current.x, current.y + 1},
				}
				for _, cd := range around {
					if cd.x < 0 || cd.x >= X || cd.y < 0 || cd.y >= Y {
						continue
					}
					if mark[cd.x][cd.y] == step-1 {
						current, step = cd, step-1
					}
				}
			}
			fmt.Println(sb.String())
			return
		}
		// 向四周探索 上左下右
		around := []coord{
			{check.x - 1, check.y},
			{check.x, check.y - 1},
			{check.x + 1, check.y},
			{check.x, check.y + 1},
		}
		arrive := mark[check.x][check.y] + 1
		for _, cd := range around {
			if cd.x < 0 || cd.x >= X || cd.y < 0 || cd.y >= Y {
				continue
			}
			if maze[cd.x][cd.y] == 1 {
				continue
			}
			// 已经访问过了
			if mark[cd.x][cd.y] != 0 {
				continue
			}
			mark[cd.x][cd.y] = arrive
			queue.PushBack(cd)
		}
	}

}

func main() {
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	breadFirst(maze, coord{0, 0}, coord{5, 4})
}
