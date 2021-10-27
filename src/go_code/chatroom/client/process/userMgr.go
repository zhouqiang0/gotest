package process

import (
	"fmt"
	"test/src/go_code/chatroom/client/model"
	"test/src/go_code/chatroom/common/message"
)

//客户端要维护的map，以及curUser,在登录成功后对它们进行初始化
var onlineUsers map[int]*message.User = make(map[int]*message.User, 20)
var CurUser model.CurUser

//处理服务器返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	fmt.Println("-------开始更新-----\t", notifyUserStatusMes.UserId)
	outputOnlineUser()
}

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历onlineUsers
	fmt.Println("当前在线用户列表(实时更新)：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}
