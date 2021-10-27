package model

//定义一个user的结构体
type User struct {
	//确定字段,为了（反）序列化成功，必须添加tag
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
