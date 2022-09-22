# Go内存模型

[官网Go内存模型](https://go-zh.org/ref/mem)

# 引言
Go内存模型也是解决了一个goroutine对某个变量的写入，如何才能被另一个读取该变量的goroutine监测到


# happens-before
在一个goroutine内部，读写操作的结果必须表现得和这些操作按照代码顺序执行一样（允许不改变程序执行结果的重排序）。
```go
a = 1
b = 2
```
上面的2行代码重排序后也不影响执行结果，但是在某一个瞬间另一个goroutine可能会发现`b = 2`，但`a != 1`的情形 

如何确保对共享变量v的读取r一定能检测到对v的写入r：
1. w `happens-before` r
2. 对v的任何其他写入操作要么`happens-before`w 要么`happens-after`r

## 哪些happens-before
1. 若包p导入包q，则q的init函数会在q的任何函数启动前完成
2. go f()语句`happens-before`f内的所有语句
3. channel: 
    * 发送总`happens-before`对应的接收完成
    * close(channel)`happens-before`从该channel接收到零值
    * 从无缓冲channel的接收`happens-before`发送完成（结合channel第一条：认为`同时`）
4. Lock: Unlock()`happens-before`下一次Lock()
5. Once: 只有一个goroutine能执行，其他goroutine会等待它执行结束: 对f()的once.Do`happens-before`其他goroutine的once.Do返回之前

