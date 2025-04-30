package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 这里就是优势了，调用决定，还不是定义时确定，比c#好
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("[%s] %d\n", from, i) // 更清晰的输出格式
		// 添加随机延迟（0~500ms）
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func main() {
	//rand.Seed(time.Now().UnixNano()) // 初始化随机种子，1.2 不需要seed了，很好
	f("direct") //串行 synchronously 执行

	go f("goroutine-1") //一个go的routinue调用，concurrently执行

	//为一个匿名函数开启一个goroutine，定义和执行同时
	go func(msg string) {
		fmt.Println(msg)
	}("going") //开启goroutine 并执行,我并没有看到交叉输出going和goroutine的情况

	go f("goroutine-2") //一个go的routinue调用，concurrently执行
	// 主 goroutine 等待足够长时间
	fmt.Println("主 goroutine 开始等待...")
	time.Sleep(10 * time.Second)
	fmt.Println("主 goroutine 结束等待")
	//一定要注意，主的goroutine一旦时间到了停止，其他的goroutine也会被关闭了。
}

/**go run .\goroutines.go
[direct] 0
[direct] 1
[direct] 2
[goroutine-1] 0
[goroutine-2] 0
going
主 goroutine 开始等待...
[goroutine-2] 1
[goroutine-1] 1
[goroutine-2] 2
[goroutine-1] 2
主 goroutine 结束等待
*/
