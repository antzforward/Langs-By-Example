package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	//语句可以位于条件之前；此语句中声明的任何变量都可以在当前分支和所有后续分支中使用。
	//说实在并不如，赋值与判断分离，除非想合并语句，类似python里面if同时声明变量。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 { //没有elif 类似关键字，等同于c++类似的
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

/**
go run .\if-else.go
7 is odd
8 is divisible by 4
either 8 or 7 are even
9 has 1 digit
*/
