package main

//URL 提供了一种统一的资源定位方式。
import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	//我们将解析这个示例 URL，其中包括方案、身份验证信息、主机、端口、路径、查询参数和查询片段。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	//明文传递信息，URL应该会改的
	fmt.Println(u.Scheme)
	//User包含所有身份验证信息；调用 Username并Password获取单独的值。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	//包含Host主机名和端口（如果存在）。用于SplitHostPort提取它们。
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	//这里我们提取了path和之后的片段#。
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	//要获取格式为字符串的查询参数k=v，请使用RawQuery。
	// 您还可以将查询参数解析为映射。解析后的查询参数映射是从字符串到字符串切片的，
	// 因此[0] 如果您只需要第一个值，请对其进行索引。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}

/**go run .\url-parsing.go
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v*/
