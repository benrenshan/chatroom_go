# chartroom_go
A small project for a simple chatroom using golang.

# 项目编码流程的设计分析

通讯逻辑示意图

![image-20220307161752440](C:\Users\syt\AppData\Roaming\Typora\typora-user-images\image-20220307161752440.png)

## 客户端部分
### 客户端登录功能
1.接收的输入id和密码

2.接收服务端的返回结果

3.判断是成功还是失败，并显示对应的页面
### 客户端发送消息流程
1.创建一个Message的结构体，结构体存放消息的类型和消息本身的数据（要有用户名和密码）

2.mes.Type = 登录的消息类型

3.mes.Data = 登录消息的内容

4.对mes和mes.Data进行序列化

5.在网络传输中，最麻烦的就是丢包问题

（1）先给服务器发送mes的长度【有多少字节】

（2）再发送消息本身

6.等待服务端发送是否合法的指令loginResMes。

## 服务端部分
### 服务端验证登录部分
1.接收用户的id密码，这里要开启协程，处理多个请求

2.验证用户的用户名和密码

3.返回结果

### 服务端接收数据的流程
1.接收客户端发送的长度以及Message的结构体的

2.接收时候要判断实际接收到的消息内容是否等于这个长度len，如果不等于就是丢包了，不相等就设置纠错协议

3.取到后对mes进行反序列化，得到原本的mes

4.再取出message.Data进行反序列化

5.取出loginMes.userId和login.userPwd

6.验证登录是否合法，返回给客户端序列化后的loginResMes（该变量用来表示是否合法）
