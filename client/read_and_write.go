package main

import (
	"chatroom_demo/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

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