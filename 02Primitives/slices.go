package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set:", s)
	fmt.Println("Get:", s[2])

	fmt.Println("len:", len(s))

	//append,增加了length 和capicity
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	//能append self吗？
	//s = append(s, s) // 不能append 一个slice, 但是按照后面copy

	//新的slice
	c := make([]string, len(s))
	copy(c, s) //array copy的效果，因为大小已经确定了
	fmt.Println("cpy:", c)

	//如果设置的大小不到s呢？如下
	//c = make([]string, len(s)-2) // 少两个
	//copy(c, s)                   //array copy的效果，因为大小已经确定了，跟我们写array copy差不多的，只拷贝到自己的范围
	//fmt.Println("cpy:", c)

	//从slice 直接 declare& assign
	l := s[2:5] //从slice[low:high]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	//这个跟numpy下面的 多维数组的表现很像，不管你多少维，我都是1维
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

/**go run .\slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
Set: [a b c]
Get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]
*/
