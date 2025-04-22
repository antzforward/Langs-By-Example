package main

import "fmt"

//返回一个函数对象，这个函数对象不用参数 返回int
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

//感觉还是不如c#的形式吧，这里类似产生了Enumerate的形式吧

func main() {
	nextInt := intSeq() //获得闭包形式

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

/**go run .\closures.go
1
2
3
1
*/
