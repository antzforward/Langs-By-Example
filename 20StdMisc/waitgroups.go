package main

import (
	"fmt"
	"sync"
	"time"
)

// 这是我们将在每个 goroutine 中运行的函数。
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	//模拟费时任务
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	//此 WaitGroup 用于等待此处启动的所有 goroutine 完成。
	// 注意：如果将 WaitGroup 显式传递给函数，则应使用指针。
	// 此处是个默认的waitgroup？
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		//启动多个 goroutine 并增加每个 goroutine 的 WaitGroup 计数器
		wg.Add(1)
		//将 Worker 调用包装在一个闭包中，以确保告知 WaitGroup 该 Worker 已完成。
		// 这样，Worker 本身就无需了解其执行过程中涉及的并发原语。
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	//阻塞直到 WaitGroup 计数器返回到 0；所有工人都会收到通知，告知他们已完成工作。
	wg.Wait()
	//请注意，此方法无法直接传播来自工作进程的错误。 这个时候就要用errgroup了。
	//https://pkg.go.dev/golang.org/x/sync/errgroup 它好像c#的asynchronous context啊

}

/**go run .\waitgroups.go
Worker 5 starting
Worker 2 starting
Worker 4 starting
Worker 1 starting
Worker 3 starting
Worker 2 done
Worker 3 done
Worker 5 done
Worker 1 done
Worker 4 done*/
