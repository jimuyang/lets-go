package main

// 设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
// 写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

// 进阶:
// 你是否可以在 O(1) 时间复杂度内完成这两种操作？

// LRU最近最少使用 实际上不需要记录使用次数，用有序链表 每次访问时都将热点数据移到表头即可

// LRUCache 最近最少使用缓存的实现
type LRUCache struct {
	cache      map[int]*LinkedNode
	capacity   int
	head, tail *LinkedNode
}

// LinkedNode 双向链表的节点
type LinkedNode struct {
	key        int
	val        int
	prev, next *LinkedNode
}

// Constructor 构造器
func Constructor(capacity int) LRUCache {
	head, tail := &LinkedNode{0, 0, nil, nil}, &LinkedNode{0, 0, nil, nil}
	head.next = tail
	tail.prev = head
	return LRUCache{make(map[int]*LinkedNode), capacity, head, tail}
}

// Get 获取值 并将key设为热点数据
func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToFirst(node)
		return node.val
	} else {
		return -1
	}
}

// Put 更新/插入 并将key设为热点数据
func (this *LRUCache) Put(key int, value int) {
	node := this.cache[key]
	if node == nil {
		// check capacity
		if len(this.cache) >= this.capacity {
			node = this.tail.prev
			delete(this.cache, node.key)
		} else {
			node = &LinkedNode{}
		}
		node.key = key
	}
	node.val = value
	this.cache[key] = node
	this.moveToFirst(node)
}

func (this *LRUCache) moveToFirst(node *LinkedNode) {
	if node.next != nil && node.prev != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	this.head.next.prev = node
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
