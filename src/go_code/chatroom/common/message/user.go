package message

//定义一个user的结构体
type User struct {
	//确定字段,为了（反）序列化成功，必须添加tag
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
	//增加一个新的字段
	UserStatus int `json:"userStatus"` //显示用户状态

	Sex string `json:"sex"`
}
