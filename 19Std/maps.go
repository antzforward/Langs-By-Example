package main

import (
	"fmt"
	"maps"
)

//就是现在语言都支持的dict[key]=value的形式，或者称之为hash表吧

func main() {
	m := make(map[string]int) //key:string,value:int
	//setter
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	fmt.Println("len:", len(m))
	//getter
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	//未设置的，value返回default值
	v3 := m["k3"]
	fmt.Println("v3:", v3)
	//getter，但是没命中会增加map吗？=》不会，会的话，就完蛋了
	fmt.Println("len:", len(m))
	//删除key
	delete(m, "k2")
	fmt.Println("map:", m)
	//全清理
	clear(m)
	fmt.Println("map:", m)

	//这里两个是value，keyIsInMap？ 防止出现默认值的情况
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

/**
go run .\maps.go
map: map[k1:7 k2:13]
len: 2
v1: 7
v3: 0
len: 2
map: map[k1:7]
map: map[]
prs: false
map: map[bar:2 foo:1]
n == n2
*/
