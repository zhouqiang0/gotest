package process

import (
	"encoding/json"
	"fmt"
	"test/src/go_code/chatroom/client/utils"
	"test/src/go_code/chatroom/common/message"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//1.创建一个mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个smsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content //群发内容
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.序列化 smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("smsProgress json.Marshal err :", err)
		return
	}
	mes.Data = string(data)

	//4.序列化mes
	sendData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("smsProgress json.Marshal 2 err :", err)
		return
	}
	//5.将sendData发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(sendData)
	if err != nil {
		fmt.Println("smsProgress tf.WritePkg err :", err)
		return
	}
	return
}
