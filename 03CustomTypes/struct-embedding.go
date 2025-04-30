package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num:%v,str:%v}\n", co.num, co.str) //直接访问，或者通过嵌入的类型名字访问，如下

	fmt.Println("also num:", co.base.num) //我们可以使用嵌入的类型名称拼出完整路径。

	fmt.Println("describe:", co.describe()) //直接调用，或者下面的转向interface调用

	type describer interface {
		describe() string
	}

	var d describer = co //嵌入带有方法的结构体可用于将接口实现赋予其他结构体。
	fmt.Println("describer:", d.describe())
}

/**
go run .\struct-embedding.go
co={num:1,str:some name}
also num: 1
describe: base with num=1
describer: base with num=1
*/
