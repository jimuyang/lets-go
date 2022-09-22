package main

// 设计并实现最不经常使用（LFU）缓存的数据结构。它应该支持以下操作：get 和 put。

// get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
// put(key, value) - 如果键不存在，请设置或插入值。当缓存达到其容量时，它应该在插入新项目之前，使最不经常使用的项目无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，最近最少使用的键将被去除。

// 进阶：
// 你是否可以在 O(1) 时间复杂度内执行两项操作？

// 链接：https://leetcode-cn.com/problems/lfu-cache
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type LFUCacheNode struct {
	key        int
	value      int
	freq       int
	next, prev *LFUCacheNode
}

type LFUCache struct {
	cache      map[int]*LFUCacheNode
	freqMap    map[int]*LFUCacheNode
	capacity   int
	head, tail *LFUCacheNode
}

func LFUConstructor(capacity int) LFUCache {
	cache := make(map[int]*LFUCacheNode)
	freqMap := make(map[int]*LFUCacheNode)
	head, tail := &LFUCacheNode{}, &LFUCacheNode{}
	head.next, tail.prev = tail, head
	return LFUCache{cache, freqMap, capacity, head, tail}
}

func (this *LFUCache) Get(key int) int {
	if cacheNode, ok := this.cache[key]; ok {
		freq := cacheNode.freq
		cacheNode.freq = freq + 1
		if this.freqMap[freq] == cacheNode {
			if cacheNode.next != this.tail && cacheNode.next.freq == freq {
				this.freqMap[freq] = cacheNode.next
			} else {
				delete(this.freqMap, freq)
			}
		} else {
			// 先排到最前面
			insertBefore(cacheNode, this.freqMap[freq])
		}

		// 将cacheNode 填入 freqMap[freq+1]
		if freq1Node, ok := this.freqMap[freq+1]; ok {
			insertBefore(cacheNode, freq1Node)
		}
		this.freqMap[freq+1] = cacheNode
		return cacheNode.value
	}
	return -1
}

func insertBefore(who, target *LFUCacheNode) {
	who.prev.next = who.next
	who.next.prev = who.prev

	who.prev = target.prev
	who.next = target

	target.prev.next = who
	target.prev = who
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity < 1 {
		return
	}

	if cacheNode, ok := this.cache[key]; ok {
		cacheNode.value = value
		this.Get(key)
		return
	}
	if len(this.cache) >= this.capacity {
		// 移除tail
		delNode := this.tail.prev
		delete(this.cache, delNode.key)
		if this.freqMap[delNode.freq] == delNode {
			delete(this.freqMap, delNode.freq)
		}
		delNode.prev.next = this.tail
		this.tail.prev = delNode.prev
	}

	newNode := &LFUCacheNode{key, value, 0, nil, nil}
	this.cache[key] = newNode
	// 先放到末尾
	this.tail.prev.next = newNode
	newNode.prev = this.tail.prev
	newNode.next = this.tail
	this.tail.prev = newNode
	// 将newNode填入freqMap[0]
	if freq0Node, ok := this.freqMap[0]; ok {
		insertBefore(newNode, freq0Node)
	}
	this.freqMap[0] = newNode
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// func main() {
// 	lfu := LFUConstructor(2)
// 	// lfu.Put(1, 1)
// 	// lfu.Put(2, 2)
// 	// fmt.Println(lfu.Get(1))
// 	// lfu.Put(3, 3)
// 	// fmt.Println(lfu.Get(2))
// 	// fmt.Println(lfu.Get(3))
// 	// lfu.Put(4, 4)
// 	// fmt.Println(lfu.Get(1))
// 	// fmt.Println(lfu.Get(3))
// 	// fmt.Println(lfu.Get(4))
// 	lfu.Put(2, 1)
// 	lfu.Put(1, 1)
// 	lfu.Put(2, 3)
// 	lfu.Put(4, 1)
// 	fmt.Println(lfu.Get(1))
// 	fmt.Println(lfu.Get(2))

// 	// lfu = LFUConstructor(0)
// 	// lfu.Put(0, 0)
// 	// fmt.Println(lfu.Get(0))

// }
