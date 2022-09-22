package main

import (
	"math/rand"
)

// 设计一个支持在平均 时间复杂度 O(1) 下， 执行以下操作的数据结构。

// 注意: 允许出现重复元素。

// insert(val)：向集合中插入元素 val。
// remove(val)：当 val 存在时，从集合中移除一个 val。
// getRandom：从现有集合中随机获取一个元素。每个元素被返回的概率应该与其在集合中的数量呈线性相关。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/insert-delete-getrandom-o1-duplicates-allowed
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type RandomCollectionNode struct {
	val   int
	index int
	next  *RandomCollectionNode
}

type RandomizedCollection struct {
	data  map[int]*RandomCollectionNode
	array []*RandomCollectionNode
}

/** Initialize your data structure here. */
func RCConstructor() RandomizedCollection {
	data := make(map[int]*RandomCollectionNode)
	array := make([]*RandomCollectionNode, 0)
	return RandomizedCollection{data, array}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	node, ok := this.data[val]
	newNode := &RandomCollectionNode{val, 0, nil}
	this.array = append(this.array, newNode)
	newNode.index = len(this.array) - 1
	if ok {
		newNode.next = node
	}
	this.data[val] = newNode
	return !ok
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	node, ok := this.data[val]
	if ok {
		if this.array[len(this.array)-1] != node {
			this.array[node.index] = this.array[len(this.array)-1]
			this.array[node.index].index = node.index
		}
		this.array = this.array[:len(this.array)-1]
		if node.next != nil {
			this.data[val] = node.next
		} else {
			delete(this.data, val)
		}
	}
	return ok
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	i := rand.Intn(len(this.array))
	return this.array[i].val
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

// func main() {
// 	rc := RCConstructor()
// 	rc.Insert(1)
// 	rc.Insert(1)
// 	rc.Insert(2)
// 	rc.Insert(1)
// 	rc.Insert(2)
// 	rc.Insert(2)

// 	rc.Remove(1)
// 	rc.Remove(2)
// 	rc.Remove(2)
// 	rc.Remove(2)

// 	fmt.Println(rc.GetRandom())
// 	fmt.Println(rc.GetRandom())
// 	fmt.Println(rc.GetRandom())
// }
