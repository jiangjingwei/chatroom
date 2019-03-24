package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"io"
	"net"
)

func main() {

	fmt.Println("服务器正在监听...")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err = ", err)
			return
		} else {
			fmt.Printf("%s 连接成功\n", conn.RemoteAddr().String())
		}

		go process(conn)

	}

}

func process(conn net.Conn) {

	defer conn.Close()

	// 循环接收客户端发送的消息
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出连接...")
				return
			} else {
				fmt.Println("readPkg err = ", err)
				return

			}

		}
		err = serverProcessMes(conn, &mes)
		if err != nil {
			fmt.Println("serverProcessMes err = ", err)
			return
		}
		fmt.Println("服务器接收到的数据是：", mes)
	}

}

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

func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginResMesType:
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		// 处理登录
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return
}

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	if loginMes.Username == "alex" && loginMes.Password == "123" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册使用"
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}
	return

}
