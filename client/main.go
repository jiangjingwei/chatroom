package main

import (
	"fmt"
	"os"
)

var (
	username string
	password string
)

func main() {

	var key int
	var loop = true

	for loop {
		fmt.Println("----------多人聊天室----------")
		fmt.Println("\t\t 1 登录	")
		fmt.Println("\t\t 2 注册	")
		fmt.Println("\t\t 3 退出	")
		fmt.Println("\t\t 请输入：	")

		// 接收用户输入选择
		fmt.Scanln(&key)

		switch key {
		case 1:
			loop = false
		case 2:
			fmt.Println("正在注册...")
			loop = false
		case 3:
			fmt.Println("已退出聊天室...")
			os.Exit(0)
		default:
			fmt.Println("输入的参数不正确，请重新输入...")
		}

		if key == 1 {
			fmt.Println("请输入用户名：")
			fmt.Scanln(&username)
			fmt.Println("请输入密码：")
			fmt.Scanln(&password)
			err := login(username, password)
			if err != nil {
				fmt.Println("用户名密码错误，请重新输入")
			} else {
				// loginMenu()
				fmt.Println("登录成功")

			}

		}

	}

}
