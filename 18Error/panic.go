package main

import "os"

//panic通常表示发生了意外错误。
//我们主要用它来快速失败，以应对正常操作中不应该发生的错误，或者我们无法妥善处理的错误。
//似乎与当时我写c++，在函数入口处加assert的含义近似，快速暴露问题
func main() {
	//运行该程序将导致其崩溃，打印错误消息和 goroutine 跟踪，并以非零状态退出。
	panic("a problem")
	//当第一个 panic in main触发时，程序会直接退出，而不会执行剩余的代码。
	// 如果你想看到程序尝试创建临时文件，请注释掉第一个 panic out
	_, err := os.Create("/tmp/file")
	if err != nil {
		//请注意，与某些使用异常来处理许多错误的语言不同，在 Go 中，尽可能使用指示错误的返回值是惯用做法。
		panic(err)
	}
}

/**go run .\panic.go
panic: a problem

goroutine 1 [running]:
main.main()
        F:/Learn/Rust-By-Example/18Error/panic.go:7 +0x25
exit status 2*/
