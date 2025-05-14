package main

//命令行参数 是参数化程序执行的常用方法。例如，go run hello.go使用方法run和 hello.go参数go。
//外面都一样的，都是xxx.exe a b c d
//区分是 1、main函数的形式 2、args 从0开始还是1开始
//要试验命令行参数，最好go build先构建一个二进制文件。
import (
	"fmt"
	"os"
)

func main() {
	//os.Args提供对原始命令行参数的访问。
	// 请注意，此切片中的第一个值是程序的路径，并os.Args[1:] 保存了程序的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

/**go build .\command-line-arguments.go
.\command-line-arguments.exe a b c d
command-line-arguments.exe a b c d]
[a b c d]
c
*/
