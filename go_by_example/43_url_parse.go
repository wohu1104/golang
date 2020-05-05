package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// 该 url 包含了一个 scheme、 认证信息、主机名、端口、路径、查询参数以及查询片段。
	s := "postgres://user:pass@host.com:12345/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println("u.scheme is ", u.Scheme)
	// User 包含了所有的认证信息， 这里调用 Username 和 Password 来获取单独的值。
	fmt.Println("u.user is ", u.User)
	fmt.Println("u.user.username is ", u.User.Username())

	p, _ := u.User.Password()
	fmt.Println("p is", p)
	fmt.Println("u. host", u.Host)
	// Host 同时包括主机名和端口信息，如果端口存在的话， 可以使用 strings.Split() 从 Host 中手动提取端口。

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host is ", host)
	fmt.Println("port is ", port)

	fmt.Println("u path is", u.Path)
	fmt.Println("u fragment is", u.Fragment)

	fmt.Println("u.rawquery", u.RawQuery)
	/*
		要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。
		你也可以将查询参数解析为一个 map。已解析的查询参数 map 以查询字符串为键，
		已解析的查询参数会从字符串映射到到字符串的切片， 因此如果您只想要第一个值，则索引为 [0]
	*/
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("m is", m)
	fmt.Println(m["k"][0])

}
