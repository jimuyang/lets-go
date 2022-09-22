package main

import "container/list"

// 现在你总共有 n 门课需要选，记为 0 到 n-1。
// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
// 给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/course-schedule
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 某门课程的依赖课程数
	depCount := make(map[int]int)
	// 某门课程被哪些课程依赖
	depMap := make(map[int][]int)
	for _, depRelation := range prerequisites {
		one, dep := depRelation[0], depRelation[1]
		depCount[one]++
		if depBy, ok := depMap[dep]; ok {
			depMap[dep] = append(depBy, one)
		} else {
			depMap[dep] = []int{one}
		}
	}
	// 没有依赖的课程可以直接学习
	var learnQueue list.List
	for i := 0; i < numCourses; i++ {
		if depCount[i] == 0 {
			learnQueue.PushBack(i)
		}
	}
	haveLearnt := 0
	// 学完课程后 依赖它的课程依赖数-1
	for learnQueue.Len() > 0 {
		learn := learnQueue.Front().Value.(int)
		learnQueue.Remove(learnQueue.Front())
		haveLearnt++
		// 依赖它的课程们
		if depBy, ok := depMap[learn]; ok {
			for _, by := range depBy {
				depCount[by]--
				if depCount[by] == 0 {
					// 可以学习这门课程
					learnQueue.PushBack(by)
				}
			}
		}
	}
	return haveLearnt == numCourses
}
