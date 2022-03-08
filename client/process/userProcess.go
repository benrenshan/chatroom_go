package process

import (
	"chatroom_demo/common/message"
	"chatroom_demo/server/utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)
type UserProcess struct {
	//封装—工厂模式

}

//Login 完成登录校验,把ID和密码传过来
func (this *UserProcess) Login(userID int, userPwd string)(err error){

	//fmt.Printf("userID = %d userPwd = %s",userID,userPwd)
	//return nil
	//连接到服务器端
	conn, err := net.Dial("tcp","localhost:8889")
	if err != nil{
		fmt.Println("net.Dial err = ",err)
		return
	}
	defer conn.Close()
	//准备通过conn发送消息给服务器
	var mes message.Message  //定义一个需要序列化的mes
	mes.Type = message.LoginResMesType //类型就是字符型常量
	//创建一个LoginMes结构体,并把ID信息和密码填进去
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd
	//将loginMes json序列化再放入mes的Data中
	data, err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("json.Marshal err=", err)
		return
	}
	//把data赋给mes.Data字段
	mes.Data = string(data)
	//将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err= ", err)
		return
	}
	//data就是我们要发送的消息
	//先把data的长度发送给服务器
	//由于conn.Write里面只能写入byte类型的数据所以要把它转成byte
	//先获取到data的长度->转成一个可以表示长度的byte切片
	//有一个包encoding/binary有一个方法是ByteOrder，可以把uint32转成
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
	//fmt.Printf("客户端发送消息的长度成功！len=%d content = %s",len(data), string(data))
	//return
	_, err = conn.Write(data)
	if err != nil{
		fmt.Println("conn.Write(data) err", err)
		return
	}
	//time.Sleep(20*time.Second)
	//fmt.Println("休眠")
	//返回的消息
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn : conn,
	}
	mes,err = tf.ReadPkg()
	if err != nil{
		fmt.Println("readPkg(conn) err=", err)
		return
	}
	//将mes的Data部分反序列化 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200{
		fmt.Println("登录成功")

		go serverProcessMes(conn)

		for{
			ShowMenu()

		}
	}else{
		fmt.Println(loginResMes.Error)
	}
	return
}