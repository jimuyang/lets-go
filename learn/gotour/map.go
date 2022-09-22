package main

import (
	"fmt"
)

type Vertex1 struct {
	Lat, Lng float64
}

func main15() {
	m := make(map[string]Vertex1)
	fmt.Println(m)

	m["Bell"] = Vertex1{1.0, 2.0}
	fmt.Println(m)

	m = map[string]Vertex1{
		"bell": {1.0, 2.0},
		"lab":  {1.0, 2.0},
	}
	fmt.Println(m)

	// 删除元素
	delete(m, "lab")

	// 双赋值来检测键是否存在
	value, exist := m["lab"]
	fmt.Println(value, exist)

	WordCount("")
}

func WordCount(s string) (result map[string]int) {
	fmt.Println(result) // map[]
	return
}
