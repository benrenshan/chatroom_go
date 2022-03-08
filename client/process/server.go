package process

import (
	"chatroom_demo/client/utils"
	"fmt"
	"net"
	"os"
)

//ShowMenu 显示登录成功后的界面
func ShowMenu(){
	fmt.Println("-------恭喜xxx登录成功---------")
	fmt.Println("-------1. 显示在线用户列表---------")
	fmt.Println("-------2. 发送消息---------")
	fmt.Println("-------3. 信息列表---------")
	fmt.Println("-------4. 退出系统---------")
	fmt.Println("请选择(1-4):")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表-")
		//outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说的什么:)")
		//fmt.Scanf("%s\n", &content)
		//smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统...")
		os.Exit(0)
	default :
		fmt.Println("你输入的选项不正确..")
	}

}
//和服务器端保持通讯
func serverProcessMes(conn net.Conn){
	//创建一个transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for{
		fmt.Println("客户端不停的读取，正在等待服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil{
			fmt.Println("ReadPkg err= ", err)
			return
		}
		//如果读取到了消息，又是下一步处理逻辑
		fmt.Printf("mes = %v\n", mes)
	}
}