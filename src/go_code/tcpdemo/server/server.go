package main

import (
	"fmt"
	"net"
)

//服务器处理协程
func process(conn net.Conn) {
	//循环接收客户端发送的数据
	defer conn.Close() //关闭conn
	for {
		//创建一个切片
		buf := make([]byte, 1024)
		//等待客户端通过conn发送消息
		//如果客户端没有write[发送]，那么协程就阻塞在这
		//fmt.Printf("服务器在等待客户端%s发送消息\n", conn.RemoteAddr())
		n, err := conn.Read(buf) //从conn读取,n表示实际读到的数目
		if err != nil {
			fmt.Println("服务器端read err:", err)
			return
		}
		//显示客户端发送的数据到服务器终端
		fmt.Print(string(buf[:n]))

	}
}

func main() {
	fmt.Println("服务器开始监听...")
	//使用tcp协议，监听8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer listen.Close()

	//循环等待客户端连接
	for {
		fmt.Printf("等待客户来连接\n")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept()err:", err)
			return
		} else {
			fmt.Printf("Accept() suc con =%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//准备一个协程，为客户端服务
		go process(conn)

	}
	//fmt.Printf("listen suc=%v\n", listen)

}
