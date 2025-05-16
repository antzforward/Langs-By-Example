package main

//context.Context来控制取消操作。
// Context可以跨 API 边界和 Goroutine 传递截止时间、取消信号以及其他请求范围内的值。
import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	//context.Context机器为每个请求创建一个 A net/http，并可以使用该Context()方法。
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")
	//windows下面直接测试，不能通过 curl要换成wget吧
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done(): //上下文的Err()方法返回一个错误，解释为什么Done()通道被关闭。
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

/**go run .\context.go
ctr+z
exit status 0xc000013a

*/
