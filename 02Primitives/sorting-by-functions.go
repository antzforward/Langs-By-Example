package main

import (
	"cmp"
	"fmt"
	"slices"
)

// 自己定义cmp函数，这个比较常见的用法
func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}

/**go run .\sorting-by-functions.go
[kiwi peach banana]
[{TJ 25} {Jax 37} {Alex 72}]*/
