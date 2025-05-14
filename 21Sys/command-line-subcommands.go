package main

//这个subcommand，真得很好啊，我没相关在Python哪里实现过，感觉这个设计非常好，适应大型整合项目
//一些命令行工具，例如go或 ，git 有许多子命令，每个子命令都有各自的一组标志。
// 例如，go build和go get是 工具的两个不同的子命令go。该flag包允许我们轻松定义具有各自标志的简单子命令。
import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//我们使用该函数声明一个子命令NewFlagSet ，然后继续定义特定于该子命令的新标志。
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")
	//第二个FlagSet 很不错
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
	//命令是该程序的第一个参数。
	switch os.Args[1] {

	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

/** go build .\command-line-subcommands.go
./command-line-subcommands.exe foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]
./command-line-subcommands.exe  bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]
*/
