package main

//用于os.Exit以给定状态立即退出。
import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("!")
	//立刻马上的执行
	os.Exit(3)
}

/**go run .\exit.go
exit status 3*/
