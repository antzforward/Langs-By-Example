package main

import (
	"fmt"
	"slices"
)

// Go 的slices包实现了对内置类型和用户定义类型的排序。我们先来看一下对内置类型的排序。
func main() {

	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}

/**go run .\slices-sorting.go
Strings: [a b c]
Ints:    [2 4 7]
Sorted:  true*/
