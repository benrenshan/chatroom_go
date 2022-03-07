package main

import (
	"chatroom_demo/common/message"
	"encoding/binary"
	_"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

// readPkg 读取数据包，传入连接，返回消息本身
func readPkg(conn net.Conn)(mes message.Message, err error){
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据。")
	_, err = conn.Read(buf[:4])
	if err != nil{
		//fmt.Println("conn.Read err= ", err)
		//err = errors.New("read pkg header error")
		return
	}
	//根据buf[:4]转成uint32类型，方法是Uint32, 为了connRead使用要做类型转换
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//根据pkgLen读取mes的内容
	n, err := conn.Read(buf[:pkgLen])

	if n != int(pkgLen) || err != nil{
		fmt.Println("conn.Read fail err=", err)
		//err = errors.New("read pkg body error")
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err =", err)
		//err = errors.New("read pkg error")
		return
	}
	return
}
//writePkg 写入管道发送数据
func writePkg(conn net.Conn, data []byte)(err error){
	//先发送一个长度给客户端
	var pkgLen uint32 //pkgLen存储长度
	pkgLen = uint32(len(data))
	var buf [4]byte //4字节 也就是 32 位
	//相当于把这个长度转成len
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n!= 4 || err != nil{
		fmt.Println("conn.Write(bytes) err", err)
		return
	}
	//发送data
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil{
		fmt.Println("conn.Write(bytes) err", err)
		return
	}
	return
}
//serverProcessLogin 登录
func serverProcessLogin(conn net.Conn, mes *message.Message)(err error){
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
	//序列化之后，执行发送data, 同样进行封装到writePkg中
	err = writePkg(conn, data)
	return
}

//serverProcessMes 写一个判断信息种类的函数
func serverProcessMes(conn net.Conn, mes *message.Message)(err error){
	switch mes.Type {
		case message.LoginResMesType:
			err = serverProcessLogin(conn, mes)

		case message.RegisterMesType:
			//
		default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

//处理和客户端的通讯
func process(conn net.Conn){
	//延时关闭
	defer conn.Close()
	//循环的读取客户端发送的信息
	for{
		mes, err := readPkg(conn)
		if err != nil{
			if err == io.EOF{
				fmt.Println("客户端退出，服务端退出")
				return
			}else {
				fmt.Println("readPkg err=", err)
				return
			}
		}
		//fmt.Println("mes =", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil{
			return
		}

	}

}

func main(){
	//提示信息
	fmt.Println("服务器再8889端口监听")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil{
		fmt.Println("net.Listen err=", err)
		return
	}
	for{
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("Accept err=", err)
			return
		}
		//一旦连接成功，则启动一个协程和客户端保持通讯。
		go process(conn)

	}
}
