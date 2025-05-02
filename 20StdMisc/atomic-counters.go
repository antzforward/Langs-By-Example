package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
*
Go 中管理状态的主要机制是通过通道进行通信。我们在工作池中看到了这一点。
不过，还有其他一些管理状态的选项。
在这里，我们将研究如何使用sync/atomic包来实现由多个 Goroutine 访问的原子计数器。
*/
func main() {

	var ops atomic.Uint64

	var wg sync.WaitGroup
	//我们将启动 50 个 goroutine，每个 goroutine 将计数器增加 1000 次。
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				//为了原子地增加计数器，我们使用Add。
				ops.Add(1)
			}

			wg.Done()
		}()
	}
	//等待所有 goroutine 完成。
	wg.Wait()
	//这里没有 goroutines 写入“ops”，但使用 Load它以原子方式读取一个值是安全的，即使其他 goroutines 正在（原子地）更新它。
	fmt.Println("ops:", ops.Load())
}

/**go run .\atomic-counters.go
ops: 50000*/
