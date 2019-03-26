package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8069]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	// var buf = make([]byte, 4*1024)

	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}
	// fmt.Println("接收的长度为", this.Buf[:4])

	// 将Buf[:4]转换成uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])

	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32

	pkgLen = uint32(len(data))

	// var buf [4]byte

	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	// 发送数据长度
	_, err = this.Conn.Write(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.Write err = ", err)
		return
	}

	// 发送数据本身
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err = ", err)
		return
	}
	fmt.Printf("发送成功， 长度是：%d, 内容是：%s", len(data), string(data))
	return
}
