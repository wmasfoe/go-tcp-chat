package main

import (
	"bufio"
	"chat-client/constant"
	"fmt"
	"net"
	"os"
	"strings"
)

// 开一个 goroutine 用来读取服务端的消息
func readFromServer(conn net.Conn) {
	for {
		remoteIp := conn.RemoteAddr().String()
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("conn read error: %v \n", err.Error())
			return
		}

		fmt.Printf("%v 给你发了一条消息: %v \n", remoteIp, string(buf[:n]))
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接服务端失败")
		panic(err.Error())
	}

	connIp := conn.LocalAddr().String()

	defer func(conn net.Conn) {
		fmt.Printf("%v 关闭客户端连接", conn.LocalAddr().String())
		conn.Close()
	}(conn)

	fmt.Printf("请输入要发送的内容(输入 %v 退出聊天，输入 %v 可开启私聊): \n", constant.EXIT_FLAG, constant.CONNECT_DEMO)

	// 用来处理其他 client 发送的消息
	go readFromServer(conn)

	for {
		//读取终端输入
		reader := bufio.NewReader(os.Stdin)
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("os.Stdin error: %v", err.Error())
		}

		connByteCount, err := conn.Write([]byte(readString))
		if err != nil {
			fmt.Printf("conn write error: %v", err.Error())
			return
		}

		fmt.Printf("%v 发送了: %v，共 %v 字节\n", connIp, readString, connByteCount)

		if strings.Trim(readString, " \r\n") == constant.EXIT_FLAG {
			return
		}
	}

}
