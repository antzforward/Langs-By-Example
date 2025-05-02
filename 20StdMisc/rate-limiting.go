package main

import (
	"fmt"
	"time"
)

// NOTE:结合限速的要求一起理解：有限资源的分配原理。
// 速率限制 是控制资源利用率和维护服务质量的重要机制。
// Go 通过 goroutines、channels 和ticker优雅地支持速率限制
func main() {
	//首先我们来了解一下基本的速率限制。
	// 假设我们想要限制传入请求的处理速度。我们会将这些请求发送到同名的通道。
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	//该limiter通道每 200 毫秒接收一次值。这是我们速率限制方案中的调节器
	limiter := time.Tick(200 * time.Millisecond)

	//通过在处理每个请求之前阻止从通道接收limiter，我们将自己限制为每 200 毫秒 1 个请求。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	//我们可能希望在速率限制方案中允许短时间的突发请求，同时保持总体速率限制不变。
	// 我们可以通过缓冲限制器通道来实现这一点。该burstyLimiter 通道最多允许 3 个事件的突发。
	burstyLimiter := make(chan time.Time, 3)
	//填满通道以表示允许突发。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//每 200 毫秒，我们都会尝试向中添加一个新值burstyLimiter，最多为 3。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	//现在模拟另外 5 个传入请求。其中前 3 个将受益于 的突发功能burstyLimiter。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/**go run .\rate-limiting.go
request 1 2025-05-02 10:13:44.4927988 +0800 CST m=+0.200730901
request 2 2025-05-02 10:13:44.6930411 +0800 CST m=+0.400970001
request 3 2025-05-02 10:13:44.8922229 +0800 CST m=+0.600148601
request 4 2025-05-02 10:13:45.0927969 +0800 CST m=+0.800719401
request 5 2025-05-02 10:13:45.2928461 +0800 CST m=+1.000765401
request 1 2025-05-02 10:13:45.2928461 +0800 CST m=+1.000765401
request 2 2025-05-02 10:13:45.2928461 +0800 CST m=+1.000765401
request 3 2025-05-02 10:13:45.2928461 +0800 CST m=+1.000765401
request 4 2025-05-02 10:13:45.4937489 +0800 CST m=+1.201665001
request 5 2025-05-02 10:13:45.6937398 +0800 CST m=+1.401652701*/
