package main

import (
	"fmt"
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

// 处理和客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()

	// 这里调用总控，创建一个实例
	processor := &Processor{
		Conn: conn,
	}

	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程关闭", err)
		return
	}
}
