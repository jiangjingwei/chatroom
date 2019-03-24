package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	var buf = make([]byte, 4*1024)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}
	fmt.Println("接收的长度为", buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32

	pkgLen = uint32(len(data))

	var buf [4]byte

	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	// 发送数据长度
	_, err = conn.Write(buf[:4])
	if err != nil {
		fmt.Println("conn.Write err = ", err)
		return
	}

	// 发送数据本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err = ", err)
		return
	}
	fmt.Printf("客户端发送成功， 长度是：%d, 内容是：%s", len(data), string(data))
	return

}
