package main

import (
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

type LinkedNode struct {
	val        int
	next, prev *LinkedNode
}

// 1. 单链表 每k个做一次反转 k >= 2
func reverseEveryK(head *LinkedNode, k int) *LinkedNode {
	fakeHead := &LinkedNode{0, head, nil}
	node := fakeHead
	for node != nil {
		node = reverseAfterK(node, k)
	}
	return fakeHead.next
}

// A -> B -> C -> D -> E ... node=A,k=3 => A -> D -> C -> B -> E return B
func reverseAfterK(node *LinkedNode, k int) *LinkedNode {
	if node == nil || node.next == nil || node.next.next == nil {
		return nil
	}
	subHead, subTail := node.next, node.next
	current := node.next.next

	for i := 1; i < k && current != nil; i++ {
		next := current.next
		current.next = subHead
		subHead = current

		current = next
	}
	node.next = subHead
	subTail.next = current
	return subTail
}

// func main() {
// 	head := &LinkedNode{1, nil, nil}
// 	node := head
// 	for i := 2; i < 11; i++ {
// 		node.next = &LinkedNode{i, nil, nil}
// 		node = node.next
// 	}
// 	printLinkedList(head)
// 	head = reverseEveryK(head, 3)
// 	printLinkedList(head)
// }

func printLinkedList(head *LinkedNode) {
	var sb strings.Builder
	for ; head != nil; head = head.next {
		sb.WriteString(strconv.Itoa(head.val))
		sb.WriteString(" -> ")
	}
	fmt.Println(sb.String())
}

// 2. 设计一个限流 满足允许5秒钟一个请求
// redis 5s超时

// 2.1 5s请求不能超过2次 漏斗限流
// Redis 4.0 提供了一个限流 Redis 模块，它叫 redis-cell
// cl.throttle

// 2.2 设计一个通用限流 类似
// cl.throttle [资源key] [初始配额] [30次] [60秒]

// Limiter 限流
type Limiter struct {
	init           int
	quota, seconds int
}

// Init 初始化
func (limiter *Limiter) Init(init, quota, seconds int) {
	limiter.init = init
	limiter.quota = quota
	limiter.seconds = seconds
}

// 3. 手上有一堆扑克牌，牌面顺序是1234(从上到下)，
// 规则是，
// （1.取手中最上面的牌放入桌上牌堆的最上面
// （2.如果手中还有牌，则将手中最上面的牌放入手中牌堆的最下面
// （3.重复1、2步骤
// 这时桌上的牌堆顺序是4231(从上到下)
// 题目：已知桌上牌堆的顺序，求原来手中牌堆的顺序

func cardForwardProcess(cards []int) {
	var hand, table list.List
	putTable := true
	for _, val := range cards {
		if putTable {
			table.PushFront(val)
		} else {
			hand.PushBack(val)
		}
		putTable = !putTable
	}
	if hand.Len() > 0 {
		for node := hand.Front(); node != nil; {
			hand.Remove(node)
			if putTable {
				table.PushFront(node.Value.(int))
			} else {
				hand.PushBack(node.Value.(int))
			}
			if hand.Len() == 1 {
				node = hand.Front()
			} else {
				node = node.Next()
			}
			putTable = !putTable
		}
	}
	printIntList(table)
}

func printIntList(l list.List) {
	var sb strings.Builder
	for node := l.Front(); node != nil; node = node.Next() {
		sb.WriteString(strconv.Itoa(node.Value.(int)))
		sb.WriteString(" ")
	}
	fmt.Println(sb.String())
}

func cardReverseProcess(cards []int) {
	var hand list.List
	for _, val := range cards {
		// 放到手中最上面
		hand.PushFront(val)
		// 手中最下面放到最上面
		back := hand.Back()
		hand.Remove(back)
		hand.PushFront(back.Value.(int))
	}
	// 多执行了一次  手中最下面放到最上面
	front := hand.Front()
	hand.Remove(front)
	hand.PushBack(front.Value.(int))

	printIntList(hand)
}

// func main() {
// 	cardForwardProcess([]int{1, 2, 3, 4})
// 	cardReverseProcess([]int{4, 2, 3, 1})
// }

// 4.有一个数组 -1 3 2 4 5 -6 7 -9,将该数组的负数和正数分离开
// 多余的数放入数组的尾部,比如3 -1 2 -6 4 -9 5 7
func splitPosNeg(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {

	}
}

// 5. leetcode 670
// 给一个非负整数 最多交换一次 返回最大数字
func maximumSwap(num int) int {
	str := strconv.Itoa(num)
	runes := []rune(str)
	ints := make([]int, len(runes))
	for i, ch := range runes {
		ints[i] = int(ch - '0')
	}
	sort.Ints(ints)
	fmt.Println(ints)

	changing := false
	wait, target, targetIndex := 0, 0, 0
	for i := len(ints) - 1; i >= 0; i-- {
		n := int(runes[len(ints)-1-i] - '0')
		if changing {
			if n == target {
				// runes[len(ints)-1-i] = rune(wait + '0')
				targetIndex = len(ints) - 1 - i
			}
		} else {
			if n != ints[i] {
				runes[len(ints)-1-i] = rune('0' + ints[i])
				changing = true
				wait, target = n, ints[i]
			}
		}
	}
	if changing {
		runes[targetIndex] = rune(wait + '0')
	}

	i, _ := strconv.Atoi(string(runes))
	return int(i)
}

// func main() {
// 	fmt.Println(maximumSwap(2736))
// 	fmt.Println(maximumSwap(9973))
// }

// 6. 一个无序的数组，求这个数组的中位数
// 排序取中间
func midTerm(nums []int) {

}

// 7. 假设是一个抽奖的游戏，不同的人是有不同的概率倍数，是一个整数，例如1、3、5...
// 输入100万人，要求抽奖抽出来2万个人；并且假设每个人都有一个唯一id
// 写一个函数做下抽奖，输入和输出的数据结构自己设计

// 8. 二叉树前序遍历的递归和非递归解法

// TreeNode
type TreeNode struct {
	val         int
	left, right *TreeNode
}

func preOrderTravel(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.val)
	preOrderTravel(node.left)
	preOrderTravel(node.right)
}

func preOrderTravel1(node *TreeNode) {
	var l list.List
	l.PushFront(node)

}

// 11.一千个棋子，甲先取乙后取，每次最多取七个最少取一个，问是否有一个方案让甲一定赢
// 1000 / 8 = 125

// 12. 3×7的格子，从左上角到右下角，只能往右或者往下，有多少种走法
// 动态规划
func countPath(m, n int) int {
	if m <= 1 || n <= 1 {
		return 1
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 首行首列
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 13.一个不均匀硬币，如何抛出均匀概率
// 正反和反正的概率一样  如果两次相同则再来

// 14.然后有一个生成0到13随机数的算法，如何用它均匀生成0到9随机数
// 一直执行 直到出现0-9为止

// 15.两千万高考生成绩如何排序

// 16 手写排序

// 选择排序
func chooseSort(nums []int) {
	// 每次选择最小的元素
	l := len(nums)
	for i := 0; i < l; i++ {
		min, minIndex := nums[i], i

		for j := i + 1; j < l; j++ {
			if nums[j] < min {
				min, minIndex = nums[j], j
			}
		}
		nums[minIndex] = nums[i]
		nums[i] = min
	}
}

// 插入排序 扑克牌排序
func insertSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		// 有序部分 nums[:i] 待插入： nums[i]
		wait := nums[i]
		for j := 0; j < i; j++ {
			if nums[j] <= wait {
				continue
			} else {
				// 正确位置j
				for k := i; k > j; k-- {
					nums[k] = nums[k-1]
				}
				nums[j] = wait
				break
			}
		}
	}
}

// 冒泡排序 选择排序的变种?
func bubblingSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		// 冒泡目标位置nums[i] 冒泡路径nums[i:]
		for j, min := l-1, nums[l-1]; j >= i; j-- {
			if nums[j] > min {
				// 下移
				nums[j+1] = nums[j]
				nums[j] = min
			} else {
				min = nums[j]
			}
		}
	}
}

// 快速排序
func quickSort(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	tag, lessTo := nums[0], 0
	for i := 1; i < l; i++ {
		if nums[i] <= tag {
			lessTo++
			nums[lessTo], nums[i] = nums[i], nums[lessTo]
		}
	}
	nums[lessTo], nums[0] = tag, nums[lessTo]
	quickSort(nums[:lessTo])
	quickSort(nums[lessTo+1:])
}

// 归并排序
// func mergeSort(nums []int) {
// 	l := len(nums)
// }

func main() {
	nums := [50]int{}
	for i := 0; i < 50; i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println(nums)
	// chooseSort(nums[:])
	// insertSort(nums[:])
	// bubblingSort(nums[:])
	// quickSort(nums[:])
	fmt.Println(nums)
}

// 后台题库

// 1. 二叉树镜像
func mirrorOfTree(node *TreeNode) {
	if node == nil {
		return
	}
	node.left, node.right = node.right, node.left
	mirrorOfTree(node.left)
	mirrorOfTree(node.right)
}

// 2. 寻找第k大的素数 数组中包含不是素数的
func primeK(nums []int) {

}

// 3. 求无向无环图的直径

// 4. 手写最小堆

// 5. 群消息已读功能
// 网上有一个思路 每个群成员只记录自己已读的最新消息 每条消息发出后 需要给每个群成员插入一条未读记录

// 6. 有一个n边形 现以P0为起点将n边形的周长分成k段，每段的长度相同 打印所有的k等分点

// 点
type Point struct {
	x, y float64
}

func calcLength(p1, p2 *Point) float64 {
	return math.Abs(p1.x-p2.x) + math.Abs(p1.y-p2.y)
}

func takePoint(start, end *Point, length float64) *Point {
	if start.x == end.x {
		if end.y > start.y {
			return &Point{start.x, start.y + length}
		}
		return &Point{start.x, start.y - length}
	} else {
		if end.x > start.x {
			return &Point{start.x + length, start.y}
		}
		return &Point{start.x - length, start.y}
	}
}

func kFen(points []*Point, k int) []*Point {
	// points = append(points, points[0])
	myPoints := make([]*Point, len(points)+1)
	for i := 0; i < len(points); i++ {
		myPoints[i] = points[i]
	}
	myPoints[len(points)] = points[0]
	points = myPoints

	// 周长
	perimeter := 0.0
	last := points[0]
	for _, p := range points {
		perimeter += calcLength(last, p)
		last = p
	}
	each := perimeter / float64(k)

	result := make([]*Point, 0)
	startI, currentLen := 0, 0.0
	// k-1 个点
	for i := 0; i < k-1; {
		segLen := calcLength(points[startI], points[startI+1])
		if segLen+currentLen > each {
			// 就在这段上
			t := takePoint(points[startI], points[startI+1], segLen+currentLen-each)
			result = append(result, t)
			i++
			points[startI] = t
			currentLen = 0.0
		} else {
			startI++
			currentLen += segLen
		}
	}
	return result
}

// 2. 单向链求和
// 1 —> 2 -> 3 -> 4 + 3 -> 4
// leetcode 445

// 二叉树的左视图
// 本质上就是广度优先遍历时每一层的第一个节点

// 实现二叉搜索树的Iterator leetcode 173

// redis
// redis的数据类型有：string hash(java Map<String,String> 适合存储对象) list(go list.List) set(HashSet<String>) sortedSet(HashMap<String, Double>)
// 都是针对value而言

// redis实现分布式锁
// 互斥性：同一时刻只有一个客户端拥有锁     原子加锁            redis.set(key[商品id] value(UUID) NX EX 10)
// 安全性：锁只能由拥有的客户端解除        原子解锁 lua脚本实现  if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end
// 死锁问题：拥有锁的客户端宕机后 锁能释放  设置超时时间
// 高可用：部分节点宕机后还能正常获取锁     必须集群部署redis

// redis高可用方案
// 主从 master-slave 可以读写分离 但需要手动晋升主节点
//     哨兵机制
// redis cluster 负载均衡 虚拟槽 16384个槽 master节点维护着bitmap来直到哪些槽属于自己 集群维护槽->cluster的映射

// redis和memcached
// redis支持持久化 rdb和aof；redis数据结构更丰富
// memcached 支持图片视屏等

// 缓存击穿: 大量不在缓存的请求入站 导致db崩溃
// 缓存雪崩: 某一时间点 大量缓存同时失效 => 设置不同的过期时间/加锁or队列的方式对访问db进行限流/多层缓存 热点数据靠前但过期时间短

// 布隆过滤器：
// 本质上是一个bitmap 使用多个hash函数 功能是判断是一个值 是否一定不存在  bitmap的实现不可删除
// 因此适合在缓存击穿情况下 去除不合理请求对存储的压力

// 限流算法：漏斗/令牌桶
// 令牌桶的缺点： 只要有令牌就允许 一个很耗资源的操作可能让整个系统很久都反应不过来

// mysql存储引擎
// mysql架构： 请求 => 连接和线程处理 => sql解析器/缓存 => sql优化器 => 存储引擎
// InnoDB: 支持聚簇索引 MVCC(多版本并发控制 2个版本号) 支持高并发 行锁 支持事务
// MyISAM: 非聚簇索引 索引与数据分离 表锁 不支持事务 专注性能

// 聚簇索引和非聚簇索引的区别  索引和数据放在一起，找到索引即找到数据 就是聚簇索引
// 因此具有唯一性 只能有一个聚簇索引
// InnoDB使用聚簇索引 MyISAM使用非聚簇索引(索引和数据分离)

// sql
// 数据库事务隔离级别 InnoDB

// mysql InnoDB的事务隔离级别
// 事务隔离要解决的问题：
// 1. 脏读：读到其他事务未提交的数据 认为是脏数据
// 2. 重复读不一致： 事务中因别的事务提交了数据导致前后2次读到的数据不一致
// 3. 幻读：为了保证事务中读到的数据一致 而在事务结束后才发现别的事务提交的数据 要解决幻读 只能上表级锁 事务串行执行

// read uncommitted 读到别人未提交的数据 会发生脏读 会发生重复读不一致
// read committed   只会读到其他事务提交的数据  解决了脏读
// repeated read    允许了重复读 内部所有查询都与事务开始时一致 应该也要归功于MVCC 解决了不可重复读
// serializable     事务串行执行             解决了幻读

// 每门课程都大于80分的学生
// 1. select name from (select name, min(score) as m_score from t_score group by name having m_score > 80) temp
// 2. select name from t_score group by name having min(score) > 80

// team表，里面只有一个字段name, 一共有4条纪录，分别是a,b,c,d, 对应四个球对，现在四个球对进行比赛 写出所有对阵
// 1. select a.name, b.name from team a, team b where a.name < b.name

// 行转列
// select o.year,
//        (select amount from year_amount i where i.year = o.year and i.month = 1) as m1,
//        (select amount from year_amount i where i.year = o.year and i.month = 2) as m2,
//        (select amount from year_amount i where i.year = o.year and i.month = 3) as m3,
//        (select amount from year_amount i where i.year = o.year and i.month = 4) as m4
// from year_amount o
// group by o.year;

// 复制表结构
// create table new_table like old_table;
// 数据复制： insert into new_table select * from old_table
// 结构数据一起： create table new_table select * from old_table

// 时间比较
// select * from schedule
// where timestampdiff(minute , clock, current_timestamp()) < 5

// b树 一种平衡多路搜索树
// B树的内部节点至少有t(最小度数)个孩子（t>=2 由分裂机制保证） 最多有2t个孩子（半满 对于B*树要求至少2/3满）

// b+ b树  多路查找树
// b+树中只有页节点存储实际数据 非叶节点只存索引 数据变小因此减少io； 叶子节点通过链表的方式连接起来 方便遍历

// select poll epoll
// select poll 没有本质区别 都是轮询fd 只是数据结构一个是数组一个是链表
// epoll(event poll) 事件驱动 O(1)的复杂度 只关心活跃的连接

// 三次握手
// C:喂 听得到吗？
// S:能听到 你能听到吗？
// C:能听到

// 四次挥手
// C:就这样吧 我没话说了 挂了吧
// S:行 等我说完你再挂吧
// S:我说完了 你挂吧
// C:那我挂了 等2MSL没声音 挂掉电话

// mq rabbitmq kafka

// 网络七层协议  3高层（应用层 表示层 会话层） TCP IP 2底层（链路层 物理层）  七层负载均衡 能理解应用协议 HTTP/s等等
// 应用层：应用程序 HTTP HTTPS telnet
// 表示层：数据格式转换
// 会话层：会话维护和管理
// 传输层：TCP（传输控制协议） UDP（用户数据协议） 协议加端口                四层负载均衡 IP+协议+端口
// 网络层：路由 IP（互联网协议） ip地址
// 数据链路层：寻址和纠错
// 物理层：物理设备

// TCP 面向连接 保证正确 3次握手 可靠 点到点 基于字节流
// UDP 不面向连接 点到点到多对多都可以 丢包 响应速度高 面向报文

// synchronized和volatile的区别
// volatile的原理是内存屏障： volatile写到volatile读之间的操作不会重排序到外面
// synchronized的原理是 monitorenter monitorexit

// redis 实现共同关注
// sadd followOfA B
// sadd followOfA C
// sadd followOfA D
// sadd followOfB C
// sadd followOfB D
// sinter followOfA followOfB

// redis缓存淘汰策略
// 到时删除 设置超时时间的同时也设置一个定时器 基本不会使用 太耗CPU
// 惰性删除 访问过期key的时候删除它
// 定期删除 每隔一段时间 删除一批过期数据
//   此时删除策略又分：抽检式删除 LRU/LFU/Random/ttl从expire字典/所有字典

//  Go生产者 消费者
func produce(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch chan int) {
	for {
		val, ok := <-ch
		if ok {
			fmt.Println("consumer consume ", val)
		} else {
			fmt.Println("chan closed")
			break
		}
	}
}

// 分布式事务
// 事务 ACID
// A Atom 原子性 事务里的操作要么全执行 要么全不执行
// C consist 一致性 状态从一个一致状态转变为另一个一致状态 也隐含要求了事务的中间状态不能被观测
// I isolation 隔离性 事务和事务之间不能互相影响
// D durability 持久性 事务对数据库的修改应永久保存
// 解决方案：
// 1. 基于XA协议的两段式提交方案 XA协议主要定义了TM（全局事务管理器）和RM（局部资源管理器）之间的接口 主流的数据库都支持XA协议
// 准备阶段 和 提交阶段
// 2. TCC方案 TCC：try准备工作（冻结余额） confirm（操作余额） cancel（解冻余额）
// confirm和cancel一定要实现幂等 因为会有重试
// 3. 基于可靠消息的最终一致性方案
// 将下单和发送消息放在同一个事务里 下游库存服务听消息并允许重试
// 举个栗子 事件和工单的创建

// 一个抽奖的游戏，不同的人是有不同的概率倍数，是一个整数，例如1、3、5...
// 输入100万人，要求抽奖抽出来2万个人；并且假设每个人都有一个唯一id
// 写一个函数做下抽奖，输入和输出的数据结构自己设计

type candidate struct {
	id    int
	multi int
}

func luckyDraw(pool []candidate) {

}

// int32 = 4byte 100万 * 4 byte = 4 000 000 byte 4MB

// HTTP1.1 对比HTTP1.0
// 缓存处理 优化的缓存策略
// keep-alive 不用每次请求都要创建连接
// 支持传递Host 从而允许一个IP多个Host
// HTTP1.1中新增了24个错误状态响应码，如409（Conflict）表示请求的资源与资源的当前状态发生冲突；410（Gone）表示服务器上的某个资源被永久性的删除
// 支持断点续传 206 Partial Content

// HTTP2.0
// 二进制协议 取代文本协议
// 多路复用 原来的keep-alive是pipeline模式 现在是并发
// header压缩
// server push

// 一次完整的http请求
// * 域名解析 浏览器DNS缓存 -> 操作系统DNS缓存 -> 本地域名服务器 -> 根域名服务器 -> com域顶级域名服务器 -> 权限域名服务器 -> IP地址 -> 写入操作系统缓存 -> 写入浏览器缓存
// * http请求 传输层：TCP对HTTP请求进行封装 加入端口号 网络层：通过IP协议广播IP数据包 并获得mac地址 链路层：IP数据包封装为MAC帧 开始握手 握手后识别帧的数据部分并返回上层

// 基本工资
// 27 * 12 = 32.4

// 年终奖
// 3个月 30-40%
// 3-6  40-50%
// > 6  20%
// < 3  不完成工作的少数人

// 加班费
// 2天 * 12 = 24 * 1.2 加班1.2倍 =

// hr计算总包基数： 44万

// 公司附近住房补贴 1K
// 额外公积金 5%

// offer时间：明天/周日
