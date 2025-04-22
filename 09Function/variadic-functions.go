package main

import "fmt"

//不定参数的函数，我都快忘光了，现代的c++也支持的，其他语言不明显
func sum(nums ...int) {
	fmt.Print(nums, " ")
	totoal := 0
	for _, num := range nums { //array 循环访问, 可以按照array的形式访问
		totoal += num
	}
	fmt.Println(totoal)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...) //形式还不同哦
}

/**
go run .\variadic-functions.go
[1 2] 3
[1 2 3] 6
[1 2 3 4 5] 15
*/
