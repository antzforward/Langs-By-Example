package main

import "fmt"

//大家跟着学用tuple了？哈哈哈  (int, int)
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals() //同时dlt&assgin，有效啊，换其他语言 两句话了
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() //不想要的都是 _，来获取
	fmt.Println(c)
}

/**
go run .\multiple-return-values.go
3
7
7
*/
