package main

import (
	"fmt"
	"iter"
	"slices"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

//All 返回一个迭代器，在 Go 中它是一个具有特殊签名的函数。
/**
迭代器函数接受另一个函数作为参数，该函数yield按照约定调用（但名称可以任意）。
它会对yield每个需要迭代的元素调用该函数，并注意yield的返回值，以防可能提前终止。
*/
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {

		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

/*
*迭代不需要底层数据结构，甚至不必是有限的！
这是一个返回斐波那契数列迭代器的函数：只要yield持续返回 ，它就会持续运行true。
*/
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	//由于List.All返回一个迭代器，我们可以在常规range循环中使用它。
	for e := range lst.All() {
		fmt.Println(e)
	}

	//像slices这样的包有很多函数可以用来处理迭代器。
	//例如，Collect可以接受任意迭代器并将其所有值收集到一个切片中。
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	//一旦循环命中break或提前返回，yield传递给迭代器的函数将返回false。
	for n := range genFib() {

		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}

/**
go run .\range-over-iterators.go
10
13
23
all: [10 13 23]
1
1
2
3
5
8*/

// ## 以上的内容实在难以理解了，一般来说我们定义了一个lambda 然后函数调用的时候，调用这个lambda，
// **但是go语言这部分就写得太难懂了，这里再对比处理一下** 这个部分在md文件总结吧
