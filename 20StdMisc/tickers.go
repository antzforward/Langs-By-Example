package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		//for + select == event mode，非常好写
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C: //当前时间
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond) //1600%500=3
	ticker.Stop()
	done <- true //主动关闭其他的goroutine，不调用就主goroutine自动关闭
	fmt.Println("Ticker stopped")
}

/**go run .\tickers.go
Tick at 2025-05-01 11:03:54.8074633 +0800 CST m=+0.500000001
Tick at 2025-05-01 11:03:55.3074674 +0800 CST m=+1.000000001
Tick at 2025-05-01 11:03:55.8074714 +0800 CST m=+1.500000001
Ticker stopped*/
