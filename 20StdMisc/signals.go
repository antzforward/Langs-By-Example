package main

//有时我们希望 Go 程序能够智能地处理Unix 信号。
// 例如，我们可能希望服务器在收到 时正常关闭SIGTERM，或者命令行工具在收到 时停止处理输入SIGINT。
// 以下是如何在 Go 中使用通道处理信号。
// signal 在很多系统里面都有应用，我印象比较深的是Qt。
// 可以把Qt当成跨组件之间的信息通知机制也可以的。
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//Go 信号通知的工作原理是os.Signal 向通道发送值。
	// 我们将创建一个通道来接收这些通知。请注意，此通道必须是缓冲的。
	sigs := make(chan os.Signal, 1)
	//signal.Notify注册给定的通道来接收指定信号的通知。
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//我们可以在主函数中从这里接收sigs，但让我们看看如何在单独的 goroutine 中完成此操作，以演示更现实的优雅关闭场景。
	//多一个控制的chan来处理这个事情。
	done := make(chan bool, 1)
	//这个 goroutine 执行一个阻塞信号接收。
	// 这个 goroutine 执行一个阻塞信号接收。当它收到一个信号时，它会将其打印出来，然后通知程序可以完成了。当它收到一个信号时，它会将其打印出来，然后通知程序可以完成了。
	go func() {

		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	//程序将在这里等待，直到收到预期的信号（如上面的 goroutine 发送值所示done），然后退出。
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

/**go run .\signals.go
awaiting signal
^C
interrupt
exiting*/
