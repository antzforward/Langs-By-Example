package main

//不同的操作系统中能够解析到相同的文件包：filepath
//python 中是用自己的separate 符号来替换，并且能够替换这些变换
//windows 上对路径上的空格有要求，容易分解成两个字符串

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func getCurrentFile() string {
	_, filename, _, ok := runtime.Caller(1) // 参数 1 表示调用者的层级（0 是当前函数）
	if !ok {
		panic("Failed to get current file")
	}
	return filename
}

func main() {
	//用join 就不用考虑文件夹的分割符号的问题
	//如果一个dir名字包括多层级dir呢？可以的！
	p := filepath.Join("dir1", "dir2", "filename")
	//在windows下面就是\这个符号了
	fmt.Println("p:", p)
	//dir1/filename:文件夹名字后面带不带/都无关系吧
	fmt.Println(filepath.Join("dir1//", "filename"))
	//"dir1/../dir1" = "dir1/"
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	//文件夹的路径，a/b/c/d.file 返回a/b/c，返回结果方便放回到join里面
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	//dir包含多层级文件夹的形式
	fmt.Println("p:", filepath.Join(filepath.Dir(p), filepath.Base(p)))
	//我们可以检查路径是否是绝对路径。反面就是相对路径
	//IsAbs 是Is AbsolutePath的简写
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))
	//自己找到的一些信息，其他语言中的内置变量或者宏 比如
	//c++和c# 中的__FILE__，Rust的file!()，python的__file__ 在Go语言中，用运行调用解析的方式吧
	//Go的runtime包提供了caller函数，可以在运行时获取当前代码的文件路径和行号。这是最接近 __file__ 的方式
	//看getCurrentFile()的定义
	currentFile := getCurrentFile()
	fmt.Println("Current file:", currentFile)
	fmt.Println("Absolute path:", filepath.Dir(currentFile))
	//从二进制文件中的路径
	exePath, err := os.Executable()
	if err == nil {
		fmt.Println("Executable path:", exePath)
		fmt.Println("Executable dir:", filepath.Dir(exePath))
	}

	filename := "config.json"
	//.后面的格式，这里返回的.json，有好处：1取除类型名字获得文件名，2在文件路径中不容易出错
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	//可以用strings的功能，获得文件名 用法strings.TrimSuffix(filename, ext)
	fmt.Println(strings.TrimSuffix(filename, ext))
	//获得相对路径RelativePath的简写
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	//输出t\file
	fmt.Println(rel)
	//这个能否用在tree的结构上啊？
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		//很明显跟base对不上，打印err信息就是..\c\t\file
		//其实是对的 rel应该返回的是a\b，在拼一下就是"a/c/t/file"
		panic(err)
	}
	fmt.Println(rel)
}

/**go run .\file-paths.go
p: dir1\dir2\filename
dir1\filename
dir1\filename
Dir(p): dir1\dir2
Base(p): filename
p: dir1\dir2\filename
false
false
Current file: F:/Learn/Rust-By-Example/21Sys/file-paths.go
Absolute path: F:\Learn\Rust-By-Example\21Sys
Executable path: E:\Temp\go-build4274436452\b001\exe\file-paths.exe
Executable dir: E:\Temp\go-build4274436452\b001\exe
.json
config
t\file
..\c\t\file*/
