package main

import (
	"fmt"
	"io"
	"net"
	"test/src/go_code/chatroom/common/message"
	process2 "test/src/go_code/chatroom/server/process"
	"test/src/go_code/chatroom/server/utils"
)

type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes函数，根据消息类型决定调用函数的类型
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsp := &process2.SmsProcess{}
		smsp.SendGroupMes(mes)
	default:

	}
	return
}

func (this *Processor) subProcess() (err error) {
	//读取客户端数据
	for {
		//读取数据包，调用readPkg函数返回Message,err
		//创建一个Transfer实例完成读包
		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭了链接，服务器端退出")
				return err
			} else {
				fmt.Println("readPkg err : ", err)
				return err
			}
		}
		//fmt.Println("mes = ", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
