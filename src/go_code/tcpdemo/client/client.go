package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.10.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Println("conn 成功", conn)

	//一，发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin)

	for {
		//从终端读取一行输入，发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
			return
		}

		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}

		_, err = conn.Write([]byte(line + "\n")) //返回n(写入的字节数)及err
		if err != nil {
			fmt.Println("conn.Write err=", err)
			return
		}

		//fmt.Printf("客户端发送了%d 字节的数据", n)
	}

}
