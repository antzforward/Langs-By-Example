package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(sem chan struct{}, idx int, done chan bool) {
	defer func() { done <- true }() // 确保发送完成信号

	sem <- struct{}{}        // 占用槽位（当缓冲满时会阻塞）
	defer func() { <-sem }() // 释放槽位

	fmt.Printf("[%d] working...", idx)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	fmt.Printf("[%d] done\n", idx)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const (
		maxConcurrent = 3 // 最大并发数
		totalJobs     = 10
		timeout       = 5 * time.Second
	)

	sem := make(chan struct{}, maxConcurrent) // 并发控制通道
	done := make(chan bool)                   // 完成通知通道

	// 启动所有 worker
	for i := 0; i < totalJobs; i++ {
		go worker(sem, i, done)
	}

	// 你的计数方法（Go 中要用 for + select）
	count := 0
	timeoutChan := time.After(timeout)

	for {
		select {
		case <-done:
			count++
			if count == totalJobs {
				fmt.Println("\n所有任务完成！")
				return
			}
		case <-timeoutChan:
			fmt.Printf("\n超时退出，已完成 %d/%d 任务\n", count, totalJobs)
			return
		}
	}
}

/**go run .\channel-synchronization.go
[1] working...[0] working...[5] working...[0] done
[3] working...[3] done
[4] working...[1] done
[7] working...[5] done
[6] working...[7] done
[8] working...[4] done
[9] working...[9] done
[2] working...[6] done
[8] done
[2] done

所有任务完成！
*/
