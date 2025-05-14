package main

//环境变量是向 Unix 程序传递配置信息 的通用机制。让我们看看如何设置、获取和列出环境变量。
//在windows下面也有
//当然基本上所有的语言都支持这个环境变量的访问
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//要设置键/值对，请使用os.Setenv。
	// 要获取键的值，请使用os.Getenv。
	// 如果环境中不存在该键，这将返回一个空字符串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	value, exists := os.LookupEnv("BigMom")
	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("No Env variable BigMom")
	}
	fmt.Println()
	//用于os.Environ列出环境中的所有键/值对。这将返回一个字符串切片，
	// 格式为KEY=value。
	// 您可以使用strings.SplitN它们来获取键和值。这里我们打印所有键。
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}

/**go run .\environment-variables.go*/
