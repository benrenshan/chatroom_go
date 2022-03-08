package main

import (
	"chatroom_demo/common/message"
	process2 "chatroom_demo/server/process"
	"chatroom_demo/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct{
	Conn net.Conn
}


func (this *Processor) process2()(err error){
	//循环的读取客户端发送的信息
	for{
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil{
			if err == io.EOF{
				fmt.Println("客户端退出，服务端退出")
				return err
			}else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		err = this.serverProcessMes(&mes)
		if err != nil{
			return err
		}

	}

}
//serverProcessMes 写一个判断信息种类的函数
func (this *Processor) serverProcessMes(mes *message.Message) (err error){
	switch mes.Type {
	case message.LoginResMesType:
		//创建UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)

	case message.RegisterMesType:
		//
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}
