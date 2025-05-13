package main

import (
	"bytes"
	"fmt"
	"regexp"
)

// Go 内置了对正则表达式的支持。以下是 Go 中一些常见的与正则表达式相关的任务的示例。
// 跟其他语言的regular很接近的用法
func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch") //Regexp对象指针

	fmt.Println(r.MatchString("peach")) //MatchString

	fmt.Println(r.FindString("peach punch")) //FindString

	fmt.Println("idx:", r.FindStringIndex("peach punch")) //FindStringIndex

	fmt.Println(r.FindStringSubmatch("peach punch")) //FindStringSubmatch

	fmt.Println(r.FindStringSubmatchIndex("peach punch")) //FindStringSubmatchIndex

	fmt.Println(r.FindAllString("peach punch pinch", -1)) //FindAllString

	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r) //r的打印自身的regrex expression

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) //replace

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper) //ReplaceAllFunc
	fmt.Println(string(out))
}

/**go run .\regular-expressions.go
true
true
peach
idx: [0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
regexp: p([a-z]+)ch
a <fruit>
a PEACH*/
