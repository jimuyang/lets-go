package queue

// Queue 队列
type Queue []int

// Enter 入队
func (q *Queue) Enter(i int) {
	*q = append(*q, i)
}

// Exit 出队
func (q *Queue) Exit() int {
	i := (*q)[0]
	*q = (*q)[1:]
	return i
}

// IsEmpty 是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// godoc -http :6060
