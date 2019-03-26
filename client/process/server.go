package process

import (
	"fmt"
	"go_code/chatroom/server/utils"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("-------恭喜xxx登录成功-------")
	fmt.Println("-------1.显示在线用户列表-------")
	fmt.Println("-------2.发送消息-------")
	fmt.Println("-------3.消息列表-------")
	fmt.Println("-------4.退出系统-------")
	fmt.Println("-------请选择：-------")

	var key int
	fmt.Scanln(&key)

	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("你选择退出了系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确")
	}

}

// 处理服务器发送来的消息
func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	for {
		fmt.Println("客户端正在登录读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err = ", err)
			return
		}
		// 读取到消息
		fmt.Println("mes=", mes)

	}

}
