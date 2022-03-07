package message
type Message struct{
	Type string `json:"type"`//消息的类型，为了保证传递的时候是小写的所以用tag
	Data string `json:"data"`//消息的内容
}
const(
	LoginMesType = " LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)
// LoginMes 发送方的消息
type LoginMes struct{
	UserID int `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	UserName string `json:"user_name"` //用户名
}

type LoginResMes struct{
	Code int `json:"code"`//返回状态码 500 表示该用户未被注册 200 表示登录成功
	Error string `json:"error"`//返回错误信息
}
type RegisterMes struct {
	//
}