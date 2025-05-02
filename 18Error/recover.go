package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

//Go使用内置函数recover从panic中恢复，可以阻止程序中止，并让其继续执行。
// recover的含义时 recove from  panic？
func main() {
	//recover 必须在defer函数中调用。当defer被调用时，recover调用将panic 捕获
	//表现为r被调用，然后在自己的print调用中显示了panic的msg
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()
	//此句不会执行，因为上一句又panic，直接停止当前的goroutine了，进入defer调用过程了。
	fmt.Println("After mayPanic()")
}
