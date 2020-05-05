package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := os.Mkdir("subDir", 0755) // 在当前工作目录下，创建一个子目录。
	check(err)
	// 创建这个临时目录后，一个好习惯是：使用 defer 删除这个目录。os.RemoveAll 会删除整个目录（类似于 rm -rf）。
	defer os.RemoveAll("subDir")

	// 创建一个空文件函数
	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subDir/file1") //

	// 创建一个有层级的目录，使用 MkdirAll 函数，并包含其父目录。 这个类似于命令 mkdir -p。
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/file4")

	// ReadDir 列出目录的内容，返回一个 os.FileInfo 类型的切片对象。
	c, err := ioutil.ReadDir("subdir/parent")
	check(err)

	for _, entry := range c {
		fmt.Println("", entry.Name(), entry.IsDir())
		/* 输出
		child true
		file2 false
		file3 false
		file4 false
		*/
	}

	// Chdir 可以修改当前工作目录，类似于 cd。
	err = os.Chdir("subdir/parent/child")
	check(err)

	// 列出当前目录内容
	c, err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("listing sundir/parent/child")

	for _, entry := range c {
		fmt.Println("", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("visiting subdir")

	/*
		我们也可以遍历一个目录及其所有子目录。 Walk 接受一个路径和回调函数，用于处理访问到的每个目录和文件。
		filepath.Walk 遍历访问到每一个目录和文件后，都会调用 visit。
	*/
	err = filepath.Walk("subdir", visit)

}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fmt.Println("", p, info.IsDir())
	return nil
}
