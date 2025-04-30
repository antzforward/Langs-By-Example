package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	//其他goroutine 发信息到chan
	go func() {
		// 添加随机延迟（0~500ms）
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) //增加随机时间
		messages <- "ping"
	}()
	//交错型 传递信息，buffer size 2 发3个msg 看看
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages) //明显等待了一下，buffer的size 含义就是满则等吧
}

/**go run .\channel-buffering.go
buffered
channel
ping
*/
