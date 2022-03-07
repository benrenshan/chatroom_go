package main

import (
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
	var loop = true
	for loop{
		fmt.Println("欢迎登录聊天系统")
		fmt.Println("1.登录")
		fmt.Println("2.注册")
		fmt.Println("3.退出")
		fmt.Println("请选择1-3")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
			//loop = false
		default:
			fmt.Println("你的输入有错误")

		}
	}
	//根据用户的输入显示新的提示信息
	if key == 1{
		//用户登录
		fmt.Println("请输入用户的ID")
		fmt.Scanln(&userID)
		fmt.Println("请输入用户的password")
		fmt.Scanln(&userPwd)
		//先把登陆的函数，写到另外一个文件，比如login.go
		login(userID,userPwd)
		//if err != nil{
		//	fmt.Println("登陆失败")
		//}else{
		//	fmt.Println("登陆成功")
		//}
	}else if key == 2{
		fmt.Println("进行用户测试")
	}
}

