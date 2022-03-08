package utils

import (
	"chatroom_demo/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)
type Transfer struct {
	//分析它应该有哪写字段
	Conn net.Conn
	Buf [8096]byte //这是传输时候，使用缓冲

}

// readPkg 读取数据包，传入连接，返回消息本身
func (this *Transfer) ReadPkg()(mes message.Message, err error){

	fmt.Println("读取客户端发送的数据。")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil{
		//fmt.Println("conn.Read err= ", err)
		//err = errors.New("read pkg header error")
		return
	}
	//根据buf[:4]转成uint32类型，方法是Uint32, 为了connRead使用要做类型转换
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	//根据pkgLen读取mes的内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])

	if n != int(pkgLen) || err != nil{
		fmt.Println("conn.Read fail err=", err)
		//err = errors.New("read pkg body error")
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err =", err)
		//err = errors.New("read pkg error")
		return
	}
	return
}
//writePkg 写入管道发送数据
func (this *Transfer) WritePkg(data []byte) (err error){
	//先发送一个长度给客户端
	var pkgLen uint32 //pkgLen存储长度
	pkgLen = uint32(len(data))
	//相当于把这个长度转成len
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	// 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n!= 4 || err != nil{
		fmt.Println("conn.Write(bytes) err", err)
		return
	}
	//发送data
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil{
		fmt.Println("conn.Write(bytes) err", err)
		return
	}
	return
}