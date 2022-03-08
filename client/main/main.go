package main

import (
	"chatroom_demo/client/process"
	"fmt"
	"os"
)

//定义两个变量，一个表示用户id,一个表示用户密码
var userID int
var userPwd string

func main(){
	//接收用户的选择
//判断是否还继续显示菜单
	var key int
	//var loop = true
	for {
		fmt.Println("欢迎登录聊天系统")
		fmt.Println("1.登录")
		fmt.Println("2.注册")
		fmt.Println("3.退出")
		fmt.Println("请选择1-3")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的ID")
			fmt.Scanln(&userID)
			fmt.Println("请输入用户的password")
			fmt.Scanln(&userPwd)
			//完成登录
			//1.创建一个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userID, userPwd)
			//loop = false
		case 2:
			fmt.Println("注册")
			//loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
			//loop = false
		default:
			fmt.Println("你的输入有错误")

		}
	}
}

