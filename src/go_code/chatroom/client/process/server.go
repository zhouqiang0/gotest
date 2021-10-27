package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"test/src/go_code/chatroom/client/utils"
	"test/src/go_code/chatroom/common/message"
)

//显示登录成功后的界面
func ShowMenu() {
	fmt.Println("------恭喜登录成功------")
	fmt.Println("------1.显示在线用户列表------")
	fmt.Println("------2.发送消息------")
	fmt.Println("------3.信息列表------")
	fmt.Println("------4.退出系统------")
	fmt.Println("请选择（1-4）：")
	var key int
	var content string

	//将SmsProgress实例，创建者switch外以便多次使用
	smsProgress := &SmsProcess{}

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("在线列表")
		outputOnlineUser()
	case 2:
		fmt.Println("请输入群发消息")
		fmt.Scanf("%s\n", &content)
		smsProgress.SendGroupMes(content)

	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("选择退出了系统")
		os.Exit(0)
	default:
		fmt.Sprintln("输入选项不正确..")
	}

}

//和服务器端保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个Transfer实例，不停的读取
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err : ", err)
			return
		}
		//fmt.Println("mes=", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//取出用户信息，状态保存到客户端onlineUsers map[int]*User中
			//反序列化mes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)

			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			//处理其他用户的群发消息
			outputGroupMes(&mes)

		default:
			fmt.Println("服务器返回一个未知消息类型")

		}
	}
}
