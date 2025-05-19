package main

//从字符串解析数字是许多程序中的基本但常见的任务
import (
	"fmt"
	"strconv" //核心库
)

// strconv.ParseType(str,精度)
// 通过ParseFloat，这64可以知道要解析多少位精度。
func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	//Atoi是进行基本十进制 int解析的便捷函数。
	//atoi 也是c++语言中必有的API
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
