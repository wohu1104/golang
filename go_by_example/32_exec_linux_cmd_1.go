package main

import (
	"fmt"
	"os/exec"
)

func main() {

	dataCmd := exec.Command("date")

	dataOut, err := dataCmd.Output()

	if err != nil {
		panic(err)
	}
	fmt.Println("date output is ", string(dataOut))

	/*
		注意，在生成命令时，我们需要提供一个明确描述命令和参数的数组，
		而不能只传递一个命令行字符串。 如果你想使用一个字符串生成一个完整的命令，那么你可以使用 bash 命令的 -c 选项：
	*/
	lsCmd := exec.Command("bash", "-c", "ls -al")
	/*
		如果用-c 那么bash 会从第一个非选项参数后面的字符串中读取命令，
		如果字符串有多个空格，第一个空格前面的字符串是要执行的命令，也就是0,后面的是参数，即1, $2….
	*/

	// lsCmd := exec.Command("ls -al") // panic: exec: "ls -al": executable file not found in $PATH
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(lsOut))

}
