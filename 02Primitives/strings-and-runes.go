package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี" //默认utf8的形式，utf8底层用byte[]来表示

	fmt.Println("Len:", len(s)) //直接返回byte[]的长度，其实没必要的吧

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) //每个byte 用16表示，一个byte 8bit ，用两个16进制表示，如FF等
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) //用RuneCount表示Utf8的Char的数量，Java或者C#中的Char应该是对应这里的Rune

	for idx, runeValue := range s { //返回的居然是byte[]的index，秒~
		fmt.Printf("%#U starts at %d\n", runeValue, idx) //%#U 即表示了Utf8形式，又表示了正常打印格式，这个还真不清楚
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:]) //range s的作用，不过每次调整index
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' { //rune 可以认为是Char，可以直接比较了=》rune可以和Char直接比较
		fmt.Println("found so sua")
	}
}
