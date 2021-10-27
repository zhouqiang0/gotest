package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"test/src/go_code/chatroom/client/utils"
	"test/src/go_code/chatroom/common/message"
)

type UserProcess struct {
}

//写一个函数，完成登录
func (this *UserProcess) Login(id int, pwd string) (err error) {
	//开始定协议
	//fmt.Printf("userId= %d userPwd= %s\n", userId, userPwd)
	//return err

	//1.连接到服务器端
	conn, err := net.Dial("tcp", "localhost:8889")
	defer conn.Close()
	if err != nil {
		fmt.Println("net.Dial err : ", err)
		return
	}
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = id
	loginMes.UserPwd = pwd

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err : ", err)
		return
	}
	//5.把data赋给mes.Data
	mes.Data = string(data)

	//6.将mes进行序列化
	sendData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err : ", err)
		return
	}
	////7.此时sandData即为我们需要发送的数据
	////7.1先发送长度给服务器。获取sandData长度，转成一个表示长度的byte切片
	//pkgLen := uint32(len(sendData))
	//var buf [4]byte
	//binary.BigEndian.PutUint32(buf[:], pkgLen) //BigEndian.PutUint32: uint32数据->[]byte
	////发送长度
	//_, err = conn.Write(buf[:])
	//if err != nil {
	//	fmt.Println("conn.Write err : ", err)
	//	return
	//}
	//
	//fmt.Printf("客户端，发送消息的长度= %d 内容是: %s", len(sendData), string(sendData))
	////发送消息本身
	//_, err = conn.Write(sendData)
	//if err != nil {
	//	fmt.Println("conn.Write err : ", err)
	//	return
	//}

	//创建一个Transfer实例用了收发消息
	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(sendData)
	if err != nil {
		fmt.Println("注册发送包 err : ", err)
		return
	} else {
		fmt.Printf("客户端，发送消息的长度= %d 内容是: %s", len(sendData), string(sendData))
	}

	//还需要处理服务器端返回的消息
	getMes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err : ", err)
		return
	}

	//将getMes.Data反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(getMes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//fmt.Println("登录成功")
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = id
		CurUser.UserStatus = message.UserOnline

		//可以显示当前在线用户列表
		fmt.Println("-------当前用户在线列表：")
		for _, v := range loginResMes.UserIds {

			if v == id { //不显示自己
				continue
			}
			fmt.Println("用户id:\t", v)
			//完成客户端的 onlineUsers的初始化
			user := &message.User{
				UserId:     id,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}

		//在这里新启动一个协程，保持和服务器端的通讯，
		// 如果服务器有数据推送，则接收并显示在终端
		go serverProcessMes(conn)
		//1.显示登录成功后的菜单..
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}

//写一个函数，完成注册
func (this *UserProcess) Register(id int, pwd string, name string) (err error) {
	//1.连接到服务器端
	conn, err := net.Dial("tcp", "localhost:8889")
	defer conn.Close()
	if err != nil {
		fmt.Println("net.Dial err : ", err)
		return
	}
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.创建LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = id
	registerMes.User.UserPwd = pwd
	registerMes.User.UserName = name

	//4.将loginMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err : ", err)
		return
	}
	//5.把data赋给mes.Data
	mes.Data = string(data)

	//6.将mes进行序列化
	sendData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err : ", err)
		return
	}
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	//利用tf将sendData发送给服务器
	err = tf.WritePkg(sendData)
	if err != nil {
		fmt.Println("注册发送包 err : ", err)
		return
	}
	//收回服务器返回的消息
	getMes, err := tf.ReadPkg() //这里的mes就是RegisterResMes
	if err != nil {
		fmt.Println("tf.ReadPkg err : ", err)
		return
	}

	//将getMes.Data反序列化成RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(getMes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return

}
