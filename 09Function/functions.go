package main

import "fmt"

//还未出现函数对象的形式~，所以算是基础情况,所有的语言都一样的形式
func plus(a int, b int) int {
	return a + b
}

//相同类型，可以合并类型，特有的
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

/**
go run .\functions.go
1+2= 3
1+2+3 = 6
*/
