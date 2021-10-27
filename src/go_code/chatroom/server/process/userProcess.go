package process

import (
	"encoding/json"
	"fmt"
	"net"
	"test/src/go_code/chatroom/common/message"
	"test/src/go_code/chatroom/server/model"
	"test/src/go_code/chatroom/server/utils"
)

type UserProcess struct {
	Conn net.Conn
	//增加一个字段，表示该Conn的归属用户
	UserId int
}

//编写一个ServerProcessLogin函数，处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.从mes 中取出mes.Data，并直接反序列化为LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err : ", err)
		return
	}
	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.再声明一个LoginResMes，完成赋值
	var loginResMes message.LoginResMes
	//到redis数据库去完成验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		//用户登录成功，将其放入UserMgr中
		//将登陆成功的用户UserId赋值给this
		this.UserId = loginMes.UserId //！重要！
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUser(loginMes.UserId)

		//将当前用户的在线情况返回给loginResMes.UsersId
		//遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}

		fmt.Printf("user:%s 登录成功", user.UserName)
	}
	//用户id为100，密码为123，认为合法
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123"{
	//	loginResMes.Code = 200
	//}else{
	//	loginResMes.Code = 500
	//	loginResMes.Error = "该用户不存在"
	//}

	//3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}
	//6.发送data,将其封装到writePkg()中
	//使用mvc,先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return err

}

//编写一个ServerProcessRegister函数，处理用户注册请求
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	//1.从mes 中取出mes.Data，并直接反序列化为LoginMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err : ", err)
		return
	}
	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterMesResType

	//2.再声明一个RegisterResMes，完成赋值
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	//3.将registerResMes序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}
	//6.发送data,将其封装到writePkg()中
	//使用mvc,先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return err

}

//编写一个通知所有在线用户的方法
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历onlineUsers,一个一个发送
	for id, up := range userMgr.onlineUsers {
		//过滤掉自己
		if id == userId {
			continue
		}
		//调用函数NotifyMeOnline,通知其他用户
		up.NotifyMeOnline(userId)

	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	//1.开始组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//2.将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("NotifyMeOnline json.Marshal err : ", err)
		return
	}
	//3.将序列化后的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)

	//4.对mes进行序列化，
	sendData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("NotifyMeOnline json.Marshal 2 err : ", err)
		return
	}
	//5.创建Transfer实例，通过服务器发送
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(sendData)
	if err != nil {
		fmt.Println("NotifyMeOnline tf.WritePkg err : ", err)
		return
	}
}
