package main

//Go 提供对 JSON 编码和解码的内置支持，包括内置和自定义数据类型。
import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// custom type
type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	//内置的数据类型，常规用法b = json.Marshal(val) 可读的时候，用string(b)
	//就是先转成binary，再转成string形式。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	//自定义数据类型，先找memory地址，然后b=json.Marshal(addr) 在string(b) 只是传入地址来而已
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	//这个byte array 非常方便了。但是非常难理解
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//我们需要提供一个变量，用于存放 JSON 包解码后的数据。
	// 该变量 map[string]interface{}将保存字符串到任意数据类型的映射。
	var dat map[string]interface{}
	//这是实际的解码和相关错误的检查。
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	//为了使用解码后的映射中的值，我们需要将它们转换为适当的类型。
	// 例如，这里我们将值转换为num所需的float64类型。
	num := dat["num"].(float64)
	fmt.Println(num)
	//访问嵌套数据需要一系列转换。
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	//我们还可以将 JSON 解码为自定义数据类型。
	// 这样做的好处是，可以为我们的程序增加额外的类型安全性，并且消除了访问解码数据时进行类型断言的需要。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	//在上面的例子中，我们总是使用字节和字符串作为数据和标准输出 JSON 表示之间的中间体。
	// 我们也可以将 JSON 编码直接流式传输到os.Writer类似 s 的文件 中os.Stdout，甚至是 HTTP 响应体中。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
	//os.Reader从类似 sos.Stdin 或 HTTP 请求主体的流式读取是通过 完成的json.Decoder。
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}

/**go run .\json.go
true
1
2.34
"gopher"
["apple","peach","pear"]
{"apple":5,"lettuce":7}
{"Page":1,"Fruits":["apple","peach","pear"]}
{"page":1,"fruits":["apple","peach","pear"]}
map[num:6.13 strs:[a b]]
6.13
a
{1 [apple peach]}
apple
{"apple":5,"lettuce":7}
{1 [apple peach]}*/
