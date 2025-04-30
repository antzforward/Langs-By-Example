package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	messages := make(chan string)
	//其他goroutine 发信息到chan
	go func() {
		// 添加随机延迟（0~500ms）
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) //增加随机时间
		messages <- "ping"
	}()
	//主goroutine（或另外其他goroutine） 获取msg，如果不是主goroutine的话，估计就直接退出了
	msg := <-messages //此处会block，直到获得消息msg
	fmt.Println(msg)

}

/**go run .\channels.go
ping
*/
