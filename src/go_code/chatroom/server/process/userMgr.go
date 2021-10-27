package process

import "fmt"

//因为UserMgr实例在服务器端有且仅有一个
//因为在很多地方都会使用到，故定义我i全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成userMgr的初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUser的添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//完成对onlineUser的删除
func (this *UserMgr) DeleteOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前所有在线用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据ID返回对应的值
func (this *UserMgr) GetOnlineUser(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d 不在线", userId)
		return
	}
	return
}
