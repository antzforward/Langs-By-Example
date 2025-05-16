package main

//有时我们的 Go 程序需要产生其他非 Go 进程。
//在Python，我用过subprocess.run的API去开启一个进程调用其他语言的进程
//这些只能都能在command line能够输入 并且正常有结果的才行
import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	dateCmd := exec.Command("date")
	//该Output方法运行命令，等待其完成并收集其标准输出。如果没有错误，dateOut将保存包含日期信息的字节。
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
	//Output如果执行命令时出现问题（例如错误路径），或者 命令运行但以非零返回代码退出，则其他方法Command将返回 。*exec.Error*exec.ExitError
	//没有参数 -x
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

/**go run .\spawning-processes.go
> date
2025年05月16日 17:11:11

command exit rc = 1
> grep hello
hello grep

panic: exit status 1

goroutine 1 [running]:
main.main()
        F:/Learn/Rust-By-Example/20StdMisc/spawning-processes.go:52 +0x54d
exit status 2*/
