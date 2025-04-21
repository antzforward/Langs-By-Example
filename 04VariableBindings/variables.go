package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int //default int
	fmt.Println(e)

	f := "apple" //declaring and initializing a varaible
	fmt.Println(f)
}

/**
go run .\variables.go
initial
1 2
true
0
apple
*/
