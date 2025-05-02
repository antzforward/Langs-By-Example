package main

/**
在上例中，我们使用了显式互斥锁 来同步多个 Goroutine 之间对共享状态的访问。
另一种选择是使用 Goroutine 和通道的内置同步功能来实现相同的效果。
这种基于通道的方法与 Go 的共享内存理念相符，即通过通信实现内存共享，
并且每个数据都只由一个 Goroutine 拥有。*/
import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 在这个例子中，我们的状态将由一个 Goroutine 拥有。
// 这将保证数据在并发访问时不会被破坏。为了读取或写入该状态，
// 其他 Goroutine 会向拥有它的 Goroutine 发送消息并接收相应的回复。
// 这些readOp和writeOp struct封装了这些请求以及拥有它的 Goroutine 的响应方式。
// ME:可以看到，自带chan，有状态 用key表示
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var readOps uint64
	var writeOps uint64
	//reads和通道writes将分别被其他 goroutine 用于发出读取和写入请求。
	reads := make(chan readOp)
	writes := make(chan writeOp)

	//这是拥有 的 Goroutine state，它与上例一样是一个 map，
	// 但现在是状态 Goroutine 的私有成员。这个 Goroutine 反复在reads和writes通道上进行选择，
	// 并在请求到达时进行响应。响应的执行方式是，首先执行请求的操作，
	// 然后在响应通道上发送一个值resp来指示操作成功（如果是 ，则发送所需的值reads）。
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	//这将启动 100 个 Goroutine，通过通道向状态拥有者发出读取请求reads。
	// 每次读取都需要构造一个readOp，将其发送到reads通道，然后通过提供的通道接收结果resp。
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	//我们也使用类似的方法开始 10 次写入。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	//让 goroutines 工作一秒钟。
	time.Sleep(time.Second)
	//最后，捕获并报告操作计数。但是我的报告非常差，官方完成的数量是80，000 我大概是56，000 差距明显
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

/**go run .\stateful-goroutines.go
readOps: 55437
writeOps: 5548*/
