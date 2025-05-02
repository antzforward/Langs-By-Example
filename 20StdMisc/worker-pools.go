package main

import (
	"fmt"
	"time"
)

// 这是工作线程，我们将运行多个并发实例。这些工作线程将在jobs通道上接收工作，
// 并将相应的结果发送到通道results。我们将每个作业休眠一秒钟，以模拟一个开销很大的任务。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	//为了充分利用我们的员工资源，我们需要向他们发送工作并收集他们的成果。
	// 我们为此设立了两个渠道。
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	//这将启动 3 个工人，最初由于尚无工作而被阻止。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	//这里我们发送 5 jobs，然后通过close该频道表示这就是我们所有的工作。
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	//最后，我们收集所有工作结果。这也确保了工作 Goroutine 已经完成。
	// 等待多个 Goroutine 的另一种方法是使用WaitGroup。
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

/**go run .\worker-pools.go
worker 3 started  job 3
worker 2 started  job 2
worker 1 started  job 1
worker 2 finished job 2
worker 2 started  job 4
worker 3 finished job 3
worker 3 started  job 5
worker 1 finished job 1
worker 2 finished job 4
worker 3 finished job 5*/
