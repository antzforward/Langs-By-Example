package main

//程序中一个常见的需求是获取自Unix 纪元以来的秒数、毫秒数或纳秒数 。以下是如何在 Go 中实现此操作的。
import (
	"fmt"
	"time"
)

func main() {
	//还是用相同的方式获得当前的时间
	now := time.Now()
	fmt.Println(now)
	//用Unix来解析当前时间情况，输出格式就是个long int形式
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())
	//转回人可读的时间形式
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}

/**go run .\epoch.go
2025-05-06 10:17:52.5285424 +0800 CST m=+0.000000001
1746497872
1746497872528
1746497872528542400
2025-05-06 10:17:52 +0800 CST
2025-05-06 10:17:52.5285424 +0800 CST*/
