package main

import "fmt"

type rect struct {
	width, height int
	callCount     int
}

// 指针调用，但是看不出来啊
func (r *rect) area() int {
	r.callCount++
	return r.width * r.height
}

/**这里并不能作为func的变化，并不能并存的。
func (r rect) area() int {
	return r.width * r.height
}
*/

// 实体调用，这个也看不出来，但是实体会产生copy
func (r rect) perim() int {
	r.callCount++ //这个信息保存在copied 实体上，所以在函数外就无用了
	return 2*r.width + 2*r.height
}

//怎么体现实体会调用时产生copy的情况呢？
func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim", r.perim())
	fmt.Println("Call Count:", r.callCount)

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
	fmt.Println("Call Count:", rp.callCount)
}

/**
go run .\methods.go
area: 50
perim 30
Call Count: 1
area:  50
perim: 30
Call Count: 2
*/
