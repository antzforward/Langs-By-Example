package main

//在程序执行过程中，我们经常需要创建一些在程序退出后不再需要的数据。
// 临时文件和目录非常适合用于此目的，因为它们不会随着时间的推移污染文件系统。
import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//创建临时文件最简单的方法是调用os.CreateTemp。它会创建一个文件并将 其打开以供读写。
	// 我们将 设置"" 为第一个参数，因此它os.CreateTemp会在操作系统的默认位置创建该文件。
	f, err := os.CreateTemp("", "sample")
	check(err)
	//显示临时文件的名称。在基于 Unix 的操作系统上，目录可能是/tmp。
	// 文件名以第二个参数指定的前缀开头os.CreateTemp，其余部分由系统自动选择，以确保并发调用始终创建不同的文件名。
	// 路径可以设置，文件名就别想了，每次获得的文件名都不同 这里表现为sample+hash值
	fmt.Println("Temp file name:", f.Name())
	//完成后清理文件。操作系统可能会在一段时间后自动清理临时文件，但明确执行此操作是一个好习惯。
	//注意，自己清理是好习惯！
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)
	//如果我们打算写入许多临时文件，我们可能更愿意创建一个临时目录。
	// os.MkdirTemp的参数与的相同 CreateTemp，但它返回目录名 而不是打开的文件。
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)
	//现在我们可以通过在临时文件名前加上临时目录来合成临时文件名。
	fname := filepath.Join(dname, "file1")
	//指定文件名的方式用来写入，只是FileMode 0666 是什么含义？
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
