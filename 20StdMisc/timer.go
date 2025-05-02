package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)
	//<-timer1.C计时器的通道被阻塞，直到C 它发送一个表示计时器已触发的值。
	<-timer1.C
	fmt.Println("Timer 1 fired")
	//如果你只是想等待，可以使用 time.Sleep。
	// 计时器可能有用的一个原因是你可以在它触发之前取消它。这里有一个例子。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	//给予timer2足够的时间射击（如果它要射击的话），以表明它实际上已经停止了
	time.Sleep(2 * time.Second)
}

/**go run .\timer.go
Timer 1 fired
Timer 2 stopped
*/
