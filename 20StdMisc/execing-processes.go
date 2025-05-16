package main

//当我们需要一个可供正在运行的 Go 进程访问的外部进程时，
// 我们会这样做。有时我们只是想用另一个（可能是非 Go 的）进程完全替换当前的 Go 进程。为此，我们将使用 Go 的 classic exec 函数实现。
import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	fmt.Println("ls Path is:", binary)

	args := []string{"ls", "-a", "-l", "-h"}
	//环境变量，也可以自己定义一个/修改一个，这是开放了全部设置
	env := os.Environ()
	//一种方式用 exec.Command的形式调用，感觉会简单一点哦
	//这里指向了具体的路径
	//请注意，Go 不提供经典的 Unixfork 函数。不过通常这不是问题，
	// 因为启动 goroutine、生成进程和执行进程涵盖了 的大多数用例fork
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
