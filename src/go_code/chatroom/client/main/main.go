package main

import (
	"fmt"
	"test/src/go_code/chatroom/client/process"
)

//定义两个变量，一个表示id,一个表示密码
var userId int
var userPwd string
var userName string

func main() {
	var key int
	for {
		fmt.Println("------欢迎登录聊天系统------")
		fmt.Println("\t\t\t 1 登录聊天室--------")
		fmt.Println("\t\t\t 2 注册用户--------")
		fmt.Println("\t\t\t 3 退出系统--------")
		fmt.Println("\t\t\t 请选择（1-3）：--------")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			//用户要登录了
			fmt.Println("请输入id(数字):")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码:")
			fmt.Scanf("%s\n", &userPwd)

			//1.创建一个UserProcess的实例,完成登录功能
			up := &process.UserProcess{}
			up.Login(userId, userPwd)

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字")
			fmt.Scanf("%s\n", &userName)
			//1.创建一个UserProcess的实例，完成注册功能
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")

		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
	//
	//if key == 1{
	//	//先把登录的函数，写到另一个文件，比如login.go
	//	//main2.login(userId, userPwd)
	//	//if err != nil{
	//	//	fmt.Println("登录失败")
	//	//}else {
	//	//	fmt.Println("登录成功")
	//	//}
	//}else if key == 2{
	//	fmt.Println("进行用户注册程序")
	//}
}
