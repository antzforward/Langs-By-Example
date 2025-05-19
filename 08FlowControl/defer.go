package main

import (
	"fmt"
	"os"
)

// Defer用于确保函数调用在程序执行的稍后阶段执行，通常用于清理目的。
// defer通常用于例如 ensure在finally其他语言中使用。
func main() {

	f := createFile("/tmp/defer.txt")
	//使用 获取文件对象后 createFile，注意这时才有文件handle
	// 我们立即使用 推迟关闭该文件closeFile。这将在封闭函数 ( main) writeFile完成后执行。
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

// 关闭文件时检查错误很重要，即使是在延迟函数中。
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		panic(err)
	}
}

/**go run .\defer.go
creating
panic: open /tmp/defer.txt: The system cannot find the path specified.

goroutine 1 [running]:
main.createFile({0x39233d, 0xe})
        F:/Learn/Rust-By-Example/18Error/defer.go:19 +0x85
main.main()
        F:/Learn/Rust-By-Example/18Error/defer.go:10 +0x2b
exit status 2*/
