package main

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

// func loginMenu() {
// 	var (
// 		key string
// 	)
// 	fmt.Println("1、显示在线用户列表")
// 	fmt.Println("2、发送信息")
// 	fmt.Println("3、信息列表")
// 	fmt.Println("4、退出系统")
// 	fmt.Println("\t\t 请输入：	")

// 	fmt.Scanln(&key)

// }

func login(username string, password string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	defer conn.Close()

	var mes message.Message

	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.Username = username
	loginMes.Password = password

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	err = writePkg(conn, data)
	if err != nil {
		fmt.Println("writePkg err = ", err)
		return
	}

	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg err = ", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
