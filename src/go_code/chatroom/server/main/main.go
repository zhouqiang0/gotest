package main

import (
	"fmt"
	"net"
	"test/src/go_code/chatroom/server/model"
	"time"
)

//func readPkg(conn net.Conn)(mes message.Message, err error)  {
//	buf := make([]byte, 8096)
//	fmt.Println("等待读取客户端发送的信息")
//
//	//在conn没有关闭的情况下，才阻塞
//	_, err = conn.Read(buf[:4])		//从conn中读取4个字节(消息长度)到buf中
//	if err != nil{
//		//err = errors.New("read pkg header error")
//		return
//	}
//	fmt.Println("读取的buf= ", buf[:4])
//
//	//根据buf[:4]转成一个 uint32类型
//	pkgLen := binary.BigEndian.Uint32(buf[:4])	//BigEndian.PutUint32: []byte -> uint32数据
//	//根据pkgLen读取消息内容
//	n, err := conn.Read(buf[:pkgLen])		//从conn中读取pkgLen个字节到buf中
//	if n != int(pkgLen) || err != nil{		//n(int)判断是否读到了这么多字节
//		err = errors.New("read pkg body error")
//		return
//	}
//	//把pkg 反序列化为 -> message.Message
//	err = json.Unmarshal(buf[:pkgLen], &mes)
//	if err != nil{
//		fmt.Println("json.Unmarshal err : ", err)
//		return
//	}
//	return
//}
//
//func writePkg(conn net.Conn, data []byte) (err error) {
//	//先发送一个长度给对方
//	pkgLen := uint32(len(data))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[:], pkgLen) //BigEndian.PutUint32: uint32数据->[]byte
//	//发送长度
//	_, err = conn.Write(buf[:])
//	if err != nil {
//		fmt.Println("conn.Write err : ", err)
//		return
//	}
//	//发送data
//	n, err := conn.Write(data)
//	if n != int(pkgLen) || err != nil {
//		fmt.Println("conn.Write err : ", err)
//		return
//	}
//	return
//
//}

////编写一个ServerProcessLogin函数，处理登录请求
//func ServerProcessLogin(conn net.Conn, mes *message.Message) (err error) {
//	//1.从mes 中取出mes.Data，并直接反序列化为LoginMes
//	var loginMes message.LoginMes
//	err = json.Unmarshal([]byte(mes.Data), &loginMes)
//	if err != nil{
//		fmt.Println("json.Unmarshal err : ", err)
//		return
//	}
//	//1.先声明一个resMes
//	var resMes message.Message
//	resMes.Type = message.LogResMesType
//
//	//2.再声明一个LoginResMes，完成赋值
//	var loginResMes message.LoginResMes
//
//	//用户id为100，密码为123，认为合法
//	if loginMes.UserId == 100 && loginMes.UserPwd == "123"{
//		loginResMes.Code = 200
//	}else{
//		loginResMes.Code = 500
//		loginResMes.Error = "该用户不存在"
//	}
//	//3.将loginResMes序列化
//	data, err := json.Marshal(loginResMes)
//	if err != nil{
//		fmt.Println("json.Marshal err: ", err)
//		return
//	}
//	//4.将data赋值给resMes
//	resMes.Data = string(data)
//
//	//5.对resMes序列化
//	data, err = json.Marshal(resMes)
//	if err != nil{
//		fmt.Println("json.Marshal err: ", err)
//		return
//	}
//	//6.发送data,将其封装到writePkg()中
//	err = writePkg(conn, data)
//	return err
//
//}

////编写一个ServerProcessMes函数，根据消息类型决定调用函数的类型
//func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
//	switch mes.Type {
//	case message.LoginMesType:
//		err = ServerProcessLogin(conn, mes)
//	case message.RegisterMesType:
//	default:
//
//	}
//	return
//}

//处理与客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()

	////读取客户端数据
	//for{
	//	//读取数据包，调用readPkg函数返回Message,err
	//	mes, err := readPkg(conn)
	//	if err != nil{
	//		if err == io.EOF{
	//			fmt.Println("客户端关闭了链接，服务器端退出")
	//			return
	//		}else {
	//			fmt.Println("readPkg err : ", err)
	//			return
	//		}
	//	}
	//	//fmt.Println("mes = ", mes)
	//	err = serverProcessMes(conn, &mes)
	//	if err != nil{
	//		return
	//	}
	//}
	processor := &Processor{
		Conn: conn,
	}
	err := processor.subProcess()
	if err != nil {
		fmt.Println("通讯协程err : ", err)
		return
	}

}

//编写函数，完成对userDao的初始化工作
func initUserDao() {
	//这里的pool为全局变量，在redis.go中
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	//服务器启动时，初始化连接池(redis.go)
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()

	//提示信息
	fmt.Println("服务器[新结构]在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.listen err = ", err)
		return
	}
	defer listen.Close()
	//监听成功，等待客户端连接
	for {
		fmt.Printf("等待客户端连接\n")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() err = ", err)
			return
		}
		//连接成功，启动一个协程和客户端通讯。。
		go process(conn)

	}

}
