package queue

import (
	"fmt"
)

func ExampleQueue_Exit() {
	q := Queue{1}
	q.Enter(2)
	q.Enter(3)
	fmt.Println(q.Exit())
	// output:
	// 1
}
