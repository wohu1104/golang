package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// 使用 NewFlagSet 函数声明一个子命令， 然后为这个子命令定义一个专用的 flag。
	// 对于不同的子命令，我们可以定义不同的 flag。
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	// 检查哪一个子命令被调用了。
	// 每个子命令，都会解析自己的 flag 并允许它访问后续的位置参数。
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
