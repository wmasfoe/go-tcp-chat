package main

import (
	"bufio"
	"chat-client/constant"
	"fmt"
	"net"
	"os"
	"strings"
)

const SERVER_NETWORK = "tcp"
const SERVER_ADDR = "127.0.0.1:8080"

// 开一个 goroutine 用来读取服务端的消息
func readFromServer(conn net.Conn) {

	/*
		TODO 这里有bug!
		当发送私聊，使用server做了一层代理，没办法拿到真实的ip；并且不能区分全局的消息和私聊的消息。
		需要设置一种报文结构 headers: {origin: 对方ip}, body: 消息内容 但是太复杂了
		所以需要每个 client 既是 client 也是 server。进行 P2P 连接。
		p2p时不能断掉根 server ，需要单独开启一个 goroutine。
		而根server专门用来显示当前在线的 ip 以及转发全局消息。
	*/
	for {
		// remoteIp := conn.RemoteAddr().String()
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("conn read error: %v \n", err.Error())
			return
		}

		strMsg := string(buf[:n])

		// 被代理的私聊，但无法获得真实ip，需要改造，方法见TODO
		if strMsg[0:13] == "[system-info]" {
			fmt.Println(string(buf[:n]))
		} else {
			fmt.Printf("你收到一条私聊，对方 ip 被代理: %v \n", string(buf[:n]))
		}

	}
}

func main() {
	conn, err := net.Dial(SERVER_NETWORK, SERVER_ADDR)
	if err != nil {
		fmt.Println("连接服务端失败")
		panic(err.Error())
	}

	connIp := conn.LocalAddr().String()

	defer func(conn net.Conn) {
		fmt.Printf("%v 关闭客户端连接", connIp)
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

		fmt.Printf("你发送了: %v，共 %v 字节\n", readString, connByteCount)

		if strings.Trim(readString, " \r\n") == constant.EXIT_FLAG {
			return
		}
	}

}
