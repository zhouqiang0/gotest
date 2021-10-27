package model

import (
	"net"
	"test/src/go_code/chatroom/common/message"
)

//在客户端很多地方会使用curUser,故写成全局变量（userMgr中）
type CurUser struct {
	Conn net.Conn
	message.User
}
