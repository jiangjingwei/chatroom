package main

import (
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/process2"
	"go_code/chatroom/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录逻辑
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return
}

func (this *Processor) process2() (err error) {
	for {

		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出连接...")
				return err
			} else {
				fmt.Println("readPkg err = ", err)
				return err

			}

		}

		err = this.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes err = ", err)
			return err
		}
		// fmt.Println("服务器接收到的数据是：", mes)
	}
}
