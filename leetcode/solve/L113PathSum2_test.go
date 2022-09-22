package main

import (
	"testing"
)

func Test_pathSum(t *testing.T) {
	input := "[5,4,8,11,null,13,4,7,2,null,null,5,1]"
	tree := buildTreeByBreadthFirst(input)
	t.Log(pathSum(tree, 22))
}
