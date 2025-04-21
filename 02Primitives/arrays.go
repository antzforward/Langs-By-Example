package main

import "fmt"

func main() {

	var a [5]int //创建了包含5个int的数组，里面都是默认值0
	fmt.Println("emp:", a)

	a[4] = 100 //索引从0开始，设置a[4]
	fmt.Println("set:", a)
	fmt.Println("get:", a[4]) //get

	fmt.Println("len:", len(a)) //从定义时就确定了

	b := [5]int{1, 2, 3, 4, 5} //声明并初始化{...}初始化列表，用逗号分开
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5} //让编译确定，但是这个设置不设置作用不大吧，不太智能的情况
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} //这个配合就有点难了，输出为idx: [100 0 0 400 500]，
	//正确理解为，序号3设置为400，它后面设置成500，第三前面的没有指定序号，就从0开始，设置100，然后就没设置了，就为默认值。
	//比较常见的首尾直接给值，其他的用idx：value的形式提供。
	fmt.Println("idx:", b)

	var twoD [2][3]int //与c++相关的通，高阶的2，低阶的3，对应c#的two[2,3]
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}

/**
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
dcl: [1 2 3 4 5]
idx: [100 0 0 400 500]
2d:  [[0 1 2] [1 2 3]]
2d:  [[1 2 3] [1 2 3]]
*/
