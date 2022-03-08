package main

import (
	"fmt"
	"net"
)


func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	//这里调用总控, 创建一个
	processor := &Processor{
		Conn : conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误=err", err)
		return
	}
}

func main() {

	//提示信息
	fmt.Println("服务器[新的结构]在8889端口监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来链接服务器
	for {
		fmt.Println("等待客户端来链接服务器.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=" ,err)
		}

		//一旦链接成功，则启动一个协程和客户端保持通讯。。
		go process(conn)
	}
}