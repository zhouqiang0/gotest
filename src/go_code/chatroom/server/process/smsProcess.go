package process

import (
	"encoding/json"
	"fmt"
	"net"
	"test/src/go_code/chatroom/common/message"
	"test/src/go_code/chatroom/server/utils"
)

type SmsProcess struct {
}

//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//反序列化mes.Data为smsMes,取出UserId
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("server smsPro json.Unmarshal err :", err)
		return
	}
	//序列化mes,进行转发
	transferData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("server smsPro json.Marshal err :", err)
		return
	}

	//遍历服务器onlineUsers map[int]*UserProcess
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachUser(transferData, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachUser(transferData []byte, conn net.Conn) {
	//创建Transfer实例进行转发
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(transferData)
	if err != nil {
		fmt.Println("服务器群发消息 err :", err)
		return
	}
}
