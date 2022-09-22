package main

import (
	"container/list"
	"errors"
)

// 给定一个二维网格和一个单词，找出该单词是否存在于网格中。

// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

// 示例:

// board =
// [
//   ['A','B','C','E'],
//   ['S','F','C','S'],
//   ['A','D','E','E']
// ]

// 给定 word = "ABCCED", 返回 true.
// 给定 word = "SEE", 返回 true.
// 给定 word = "ABCB", 返回 false.

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/word-search

func exist(board [][]byte, word string) bool {
	// 标准的回溯试试？
	if len(word) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	mark := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		mark[i] = make([]bool, len(board[0]))
	}

	// 栈来记录
	stack := list.New()

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// 找到首字母
			if board[i][j] == word[0] {
				if (len(word)) == 1 {
					return true
				}
				mark[i][j] = true
				stack.PushBack(coord{i, j})
				// 开始向相邻探索 按照 上右下左
				more, next := true, false
				for {
					c := stack.Back().Value.(coord)
					if more {
						more = false
						stack.PushBack(coord{c.i - 1, c.j})
						continue
					}
					if next {
						next = false
						stack.Remove(stack.Back())
						if stack.Len() == 0 {
							mark[i][j] = false
							break
						} else {
							c1 := stack.Back().Value.(coord)
							nc, err := nextCoord(c, c1)
							if err != nil {
								mark[c1.i][c1.j] = false
								next = true
							} else {
								stack.PushBack(nc)
							}
						}
						continue
					}
					coordInvalid := c.i < 0 || c.j < 0 || c.i >= len(board) || c.j >= len(board[0])
					if coordInvalid || mark[c.i][c.j] || board[c.i][c.j] != word[stack.Len()-1] {
						next = true
					} else {
						if stack.Len() == len(word) {
							return true
						}
						mark[c.i][c.j] = true
						more = true
					}
				}
			}
		}
	}
	return false
}

type coord struct {
	i int
	j int
}

func nextCoord(c coord, center coord) (coord, error) {
	switch {
	case c.j < center.j:
		// 左
		return coord{}, errors.New("")
	case c.i < center.i:
		// 上
		return coord{center.i, center.j + 1}, nil
	case c.j > center.j:
		// 右
		return coord{center.i + 1, center.j}, nil
	case c.i > center.i:
		// 下
		return coord{center.i, center.j - 1}, nil
	default:
		return coord{}, errors.New("不应该")
	}
}

// func main() {

// 	fmt.Println(exist([][]byte{
// 		{'C', 'A', 'A'},
// 		{'A', 'A', 'A'},
// 		{'B', 'C', 'D'}}, "AAB"))

// }
