package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// 使用 Join 来构建可移植(跨操作系统)的路径。 它接收任意数量的参数，并参照传入顺序构造一个对应层次结构的路径。
	// 类似 python 的 os.path.join()
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p is ", p)

	// 应该总是使用 Join 代替手动拼接 / 和 \。 除了可移植性，Join 还会删除多余的分隔符和目录，使得路径更加规范。
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir 和 Base 可以被用于分割路径中的目录和文件。 此外，Split 可以一次调用返回上面两个函数的结果。
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Dir(p):", filepath.Base(p))

	// 判断路径是否为绝对路径。
	fmt.Println(filepath.IsAbs("/dir/file"))
	fmt.Println(filepath.IsAbs("dir/file"))

	fileName := "config.json"
	ext := filepath.Ext(fileName) // 用 Ext 将扩展名分割出来。
	fmt.Println("ext is ", ext)

	fmt.Println(strings.TrimSuffix(fileName, ext)) // strings.TrmSuffix 清除扩展名之后的文件名

	// Rel 寻找 basepath 与 targpath 之间的相对路径。 如果相对路径不存在，则返回错误。
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) // ret: t/file

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) // ret: ../c/t/file
}
