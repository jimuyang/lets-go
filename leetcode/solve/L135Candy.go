package main

import (
	"container/list"
	"fmt"
)

// 老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

// 你需要按照以下要求，帮助老师给这些孩子分发糖果：

// 每个孩子至少分配到 1 个糖果。
// 相邻的孩子中，评分高的孩子必须获得更多的糖果。
// 那么这样下来，老师至少需要准备多少颗糖果呢？

// 示例 1:

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/candy
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	if len(ratings) == 1 {
		return 1
	}
	candy := make([]int, len(ratings))
	candy[0] = 1
	// 单调减的栈 fixme 这个栈可以不用
	var stack list.List
	for i := 1; i < len(ratings); i++ {
		if ratings[i] < ratings[i-1] {
			// 下降 入栈
			stack.PushFront(i - 1)
		} else if ratings[i] >= ratings[i-1] {
			if stack.Len() > 0 {
				// 栈不为空 之前在下降
				c := 1
				candy[i-1] = c
				for stack.Len() > 0 {
					elem := stack.Front()
					v := elem.Value.(int)
					stack.Remove(elem)
					c++
					if candy[v] < c {
						candy[v] = c
					}
				}
			}
			if ratings[i] == ratings[i-1] {
				candy[i] = 1
			} else {
				// 此时ratings[i-1]已确定
				candy[i] = candy[i-1] + 1
			}
		}
	}
	if stack.Len() > 0 {
		// 栈不为空
		c := 1
		candy[len(ratings)-1] = c
		for stack.Len() > 0 {
			elem := stack.Front()
			v := elem.Value.(int)
			stack.Remove(elem)
			c++
			if candy[v] < c {
				candy[v] = c
			}
		}
	}
	fmt.Println(candy)
	result := 0
	for _, v := range candy {
		result += v
	}
	return result
}

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
}
