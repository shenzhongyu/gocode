package main

import (
	"net"
	"log"
	"fmt"
	"reflect"
	"io"
)

func main () {
	addr := "www.baidu.com:80"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("访问公网IP是：", conn.RemoteAddr().String())
	fmt.Printf("客户端连接的地址及端口是：%v", conn.LocalAddr())

	fmt.Println("conn.LocalAddr()所对应的数据类型：", reflect.TypeOf(conn.LocalAddr()))
	fmt.Println("conn.RemoteAddr().String()所对应的数据类型：", reflect.TypeOf(conn.RemoteAddr().String()))

	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("向服务器发送的数据大小是:", n)
	buf := make([]byte, 1024)

	n, err = conn.Read(buf)
	// io.EOF在网络编程中表示对端把连接关闭了。
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))
	conn.Close()

}