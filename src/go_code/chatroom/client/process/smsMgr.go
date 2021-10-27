package process

import (
	"encoding/json"
	"fmt"
	"test/src/go_code/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) {
	//反序列化mes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("smsMgr json.Unmarshal err :", err)
		return
	}
	fmt.Printf("\n用户  %d说:\t%s\n", smsMes.UserId, smsMes.Content)
}
