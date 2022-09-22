package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
func recWalk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	recWalk(t.Left, ch)
	ch <- t.Value
	recWalk(t.Right, ch)
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *Tree, ch chan int) {
	recWalk(t, ch)
	close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, open1 := <-ch1
		v2, open2 := <-ch2

		if open1 && open2 {
			if v1 != v2 {
				return false
			}
		} else if !open1 && !open2 {
			return true
		} else {
			return false
		}
	}
}

func main18() {
	fmt.Println(Same(New(1), New(1)))
	fmt.Println(Same(New(1), New(2)))
}
