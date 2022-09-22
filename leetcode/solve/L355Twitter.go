package main

import "fmt"

// 设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：

// postTweet(userId, tweetId): 创建一条新的推文
// getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
// follow(followerId, followeeId): 关注一个用户
// unfollow(followerId, followeeId): 取消关注一个用户

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/design-twitter
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type TwitterUser struct {
	id        int
	followee  map[int]*TwitterUser
	lastTweet *Tweet
}

type Tweet struct {
	id     int
	userID int
	next   *Tweet
}

func (t *Tweet) getValue() int {
	return -t.id
}

type Twitter struct {
	users map[int]*TwitterUser
}

/** Initialize your data structure here. */
func TwitterConstructor() Twitter {
	return Twitter{make(map[int]*TwitterUser)}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	user := this.getUser(userId)
	newTweet := &Tweet{tweetId, userId, user.lastTweet}
	user.lastTweet = newTweet
}

func (this *Twitter) getUser(userId int) *TwitterUser {
	user, ok := this.users[userId]
	if !ok {
		user = &TwitterUser{userId, make(map[int]*TwitterUser), nil}
		this.users[userId] = user
	}
	return user
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	// 本质上就是K个有序链表合并
	user := this.getUser(userId)
	minHeap := newMinHeap(len(user.followee) + 1)
	if user.lastTweet != nil {
		minHeap.add(user.lastTweet)
	}
	for _, fuser := range user.followee {
		if fuser.lastTweet != nil {
			minHeap.add(fuser.lastTweet)
		}
	}
	result := make([]int, 0)
	for i := 0; minHeap.Size() > 0 && i < 10; i++ {
		min, _ := minHeap.takeMin()
		tweet := min.(*Tweet)
		result = append(result, tweet.id)
		if tweet.next != nil {
			minHeap.add(tweet.next)
		}
	}
	return result
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	follower, followee := this.getUser(followerId), this.getUser(followeeId)
	follower.followee[followeeId] = followee
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	follower := this.getUser(followerId)
	delete(follower.followee, followeeId)
}

// func main() {
// 	t := TwitterConstructor()
// 	t.PostTweet(1, 5)
// 	t.GetNewsFeed(1)
// }

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// MinHeap 最小堆
type MinHeap struct {
	heap     []Interface
	size     int
	capacity int
}

type Interface interface {
	getValue() int
}

func newMinHeap(cap int) *MinHeap {
	heap := make([]Interface, cap+1)
	return &MinHeap{heap, 0, cap}
}

func (me *MinHeap) min() (Interface, error) {
	if me.size <= 0 {
		return nil, fmt.Errorf("heap is empty")
	}
	return me.heap[1], nil
}

func (me *MinHeap) Size() int {
	return me.size
}

func (me *MinHeap) takeMin() (Interface, error) {
	if me.size <= 0 {
		return nil, fmt.Errorf("heap is empty")
	}
	min := me.heap[1]
	me.heap[1] = me.heap[me.size]
	me.size--
	me.heapify(1)
	return min, nil
}

func (me *MinHeap) add(val Interface) error {
	if me.size >= me.capacity {
		return fmt.Errorf("heap is full")
	}
	me.size++
	me.heap[me.size] = val
	me.bubble(me.size)
	return nil
}

// 自底而上 冒泡
func (me *MinHeap) bubble(i int) {
	for i > 1 && me.heap[i/2].getValue() > me.heap[i].getValue() {
		me.heap[i/2], me.heap[i] = me.heap[i], me.heap[i/2]
		i = i / 2
	}
}

// 自顶而下
func (me *MinHeap) heapify(i int) {
	min := i
	// parent和left的较小值
	if 2*i <= me.size && me.heap[2*i].getValue() < me.heap[i].getValue() {
		min = 2 * i
	}
	if 2*i+1 <= me.size && me.heap[2*i+1].getValue() < me.heap[min].getValue() {
		min = 2*i + 1
	}
	if min != i {
		// 需要调整
		me.heap[min], me.heap[i] = me.heap[i], me.heap[min]
		me.heapify(min)
	}
}
