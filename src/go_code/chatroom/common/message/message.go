package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LogResMes"
	RegisterMesType         = "RegisterMes"
	RegisterMesResType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

//定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffLine
	UserBusyStat
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息数据
}

//定义两个消息
type LoginMes struct { //登录消息
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct { //登录返回消息
	Code int `json:"code"` //500:未注册，200：登录成功

	//增加一个字段，保存用户id的切片
	UserIds []int

	Error string `json:"error"` //返回的错误信息
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //400:已占用，200：注册成功
	Error string `json:"error"` //返回的错误信息
}

//为配合服务器端推送用户状态变化的消息，增加一个消息类型
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

//增加一个SmsMes
type SmsMes struct {
	Content string `json:"content"`
	User           //匿名结构体message.User，相当于继承
}
