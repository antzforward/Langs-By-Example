package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	//定义worker，不断接收job
	go func() {
		for {
			//当明确的调用了close(jobs),这个时候，more应该是false了
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	//	Closing a channel indicates that no more values will be sent on it.
	// This can be useful to communicate completion to the channel’s receivers.
	close(jobs)
	fmt.Println("sent all jobs")
	//这里使用了同步方法等待worker，此时会将所有的job都recieved
	<-done
	//此后再去获取任务，ok就为false了。整体这个job系统非常简洁。
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}

/**go run .\channel-closing.go
sent job 1
sent job 2
sent job 3
sent all jobs
received job 1
received job 2
received job 3
received all jobs
received more jobs: false*/
