package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"test/src/go_code/chatroom/common/message"
)

//将这些方法关联到结构体中
type Transfer struct {
	//分析传输者需要哪些字段
	Conn net.Conn
	Buf  [8096]byte //传输时使用的缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("等待读取客户端发送的信息")

	//在conn没有关闭的情况下，才阻塞
	_, err = this.Conn.Read(this.Buf[:4]) //从conn中读取4个字节(消息长度)到buf中
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}
	fmt.Println("读取的buf= ", this.Buf[:4])

	//根据buf[:4]转成一个 uint32类型
	pkgLen := binary.BigEndian.Uint32(this.Buf[:4]) //BigEndian.PutUint32: []byte -> uint32数据
	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen]) //从conn中读取pkgLen个字节到buf中
	if n != int(pkgLen) || err != nil {         //n(int)判断是否读到了这么多字节
		err = errors.New("read pkg body error")
		return
	}
	//把pkg 反序列化为 -> message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err : ", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	pkgLen := uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen) //BigEndian.PutUint32: uint32数据->[]byte
	//发送长度
	_, err = this.Conn.Write(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.Write err : ", err)
		return
	}
	//发送data
	n, err := this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write err : ", err)
		return
	}
	return

}
