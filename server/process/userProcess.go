package process2

import (
	"chatroom_demo/common/message"
	"encoding/json"
	"fmt"
	"net"
)

import (
	"chatroom_demo/server/utils"
)

type UserProcess struct{
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message)(err error){
	//从mes中取出mes.Data，然后反序列化成LoginMes
	var LoginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&LoginMes)
	if err != nil{
		fmt.Println("json.Unmarshal fail err= ", err)
		return
	}
	//定义一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//定义一个LoginResMes
	var loginResMes message.LoginResMes
	//判断用户名的id和密码是否匹配，匹配则返回正确
	if LoginMes.UserID == 100 && LoginMes.UserPwd == "aa"{
		loginResMes.Code = 200
	}else {
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在"
	}
	//把loginResMes序列化
	data,err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("json.Marshal fail ",err)
		return
	}
	resMes.Data = string(data)
	//把resMes序列化
	data,err = json.Marshal(resMes)
	if err != nil{
		fmt.Println("json.Marshal fail ",err)
		return
	}
	//将data赋值给resMes
	resMes.Data = string(data)
	//把resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil{
		fmt.Println("json Marshal fail", err)
		return
	}
	//序列化之后，执行发送data, 同样进行封装到writePkg中
	//因为MVC先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
