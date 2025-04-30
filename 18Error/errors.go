package main

import (
	"errors"
	"fmt"
)

// 现在返回tuple最好的用法就是value，error 同时返回的情况了
func f(age int) (int, error) {
	if age == 42 {
		return -1, errors.New("can't work with 42")
	}
	return age + 3, nil
}

var ErrOutofTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutofTea
	} else if arg == 4 {
		return fmt.Errorf("making tea:%w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		//if语句同时获得，且检查，这个是优点了
		if r, e := f(i); e != nil { //e一定要获取，否则直接就panic了
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutofTea) { //errors.Is 和errors.As的情况，用来进行Error比较
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}
}

/**go run .\errors.go
f worked: 10
f failed: can't work with 42
Tea is ready!
Tea is ready!
We should buy new tea!
Tea is ready!
Now it is dark.
*/
