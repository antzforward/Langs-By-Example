package main

//这个在Python用的是argparse，感觉跟go用法差不多，当然python特点还是有得
/**
import argparse

# 创建解析器
parser = argparse.ArgumentParser(description="这是一个示例程序")

# 添加参数
parser.add_argument("input", help="输入文件路径")  # 位置参数（必需）
parser.add_argument("-o", "--output", help="输出文件路径", default="output.txt")  # 可选参数
parser.add_argument("--verbose", action="store_true", help="是否显示详细信息")  # 布尔参数

# 解析参数
args = parser.parse_args()

# 使用参数
print(f"输入文件: {args.input}")
print(f"输出文件: {args.output}")
if args.verbose:
    print("详细信息已开启")
*/

import (
	"flag" //Go 提供了一个flag支持基本命令行标志解析的包。我们将使用这个包来实现示例命令行程序。
	"fmt"
)

func main() {
	//基本标志声明适用于字符串、整数和布尔值选项。这里我们声明了一个字符串标志，
	// word并带有默认值"foo" 和简短描述。此flag.String函数返回一个字符串指针（而非字符串值）；
	// 我们将在下面看到如何使用此指针。
	//使用方式1 declare & assign
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	//也可以声明一个选项，使用程序中其他位置声明的现有变量。注意，我们需要传入一个指向标志声明函数的指针。
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	//全部设置完，开始Parse（解析参数）
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// 特别注意：command-line-flags 在windows下内置有，一定要路径严格一点
//可以缺少参数，因为所有得flags都设置了默认值，但是不能多或者错参数。自动会显示提示信息（python要主动打印）
/**go build .\command-line-flags.go
./command-line-flags.exe -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []
./command-line-flags.exe -wat
flag provided but not defined: -wat
Usage of F:\Learn\Rust-By-Example\21Sys\command-line-flags.exe:
  -fork
        a bool
  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
        a string (default "foo")
./command-line-flags.exe -h
Usage of F:\Learn\Rust-By-Example\21Sys\command-line-flags.exe:
  -fork
        a bool
  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
        a string (default "foo")
*/
