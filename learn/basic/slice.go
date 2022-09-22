package main

import "fmt"

func learnSlice() {
	arr := [...]int{1, 2, 3, 5: 1}
	fmt.Println(arr)

	s := arr[:]
	s = append(s, 9)
	fmt.Println(s)
	fmt.Println(arr)
}

func main1() {
	// learnSlice()
	learnAssign()
}

type mystruct struct {
	str string
	i   int
}

func learnAssign() {
	s1 := mystruct{"1", 1}
	s2 := s1
	s2.i = 2
	fmt.Println(s1, s2)
}
