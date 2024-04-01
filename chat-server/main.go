package main

import (
	"chat-server/constant"
	"chat-server/log"
	"net"
	"regexp"
	"strings"
)

type IPAddrMap = map[string]net.Conn

func progress(conn net.Conn, ipMap *IPAddrMap) {

	connIp := conn.RemoteAddr().String()

	defer func() {
		log.Logger.Info(connIp, "已经下线")
		// 连接退出时，从ip映射表中删除对应ip
		delete(*ipMap, connIp)
		conn.Close()
	}()

	for {
		//读取客户端输入的内容
		msgBuf := make([]byte, 1024)
		readByteCount, err := conn.Read(msgBuf)
		if err != nil {
			log.Logger.Error("客户端连接发生异常", err.Error())
			return
		}
		//将内容从字节转成string
		msg := string(msgBuf[:readByteCount])
		trimmedMsg := strings.Trim(msg, " \r\n")
		//收到退出指令
		if trimmedMsg == constant.EXIT_FLAG {
			return
		}

		isConnIp := trimmedMsg[0:2] == constant.CONNECT_IP_OPT
		// 连接其他ip，同时私有连接不会在全局展示
		if isConnIp {
			//正则匹配 ip 端口号 消息
			re := regexp.MustCompile(`->(\d+\.\d+\.\d+\.\d+):(\d+):(.*)`)
			//匹配结果
			match := re.FindStringSubmatch(trimmedMsg)
			//连接的ip+端口号
			connIp := match[1] + ":" + match[2]
			//要发送的消息
			sendMsg := match[3]
			// 从 ip 表中取connect实例
			optConn := (*ipMap)[connIp]
			//给对应的ip发送消息
			_, err := optConn.Write([]byte(sendMsg))
			if err != nil {
				log.Logger.Warn("发送消息失败", err.Error())
				return
			}
		} else {
			// 对消息做处理
			log.Logger.GlobalLog(connIp, msg)
		}
	}
}

func main() {

	listen, err := net.Listen("tcp", "0.0.0.0:8080")

	ipAndConnMap := make(IPAddrMap)

	if err != nil {
		panic(err.Error())
	}

	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			panic(err.Error())
		}
	}(listen)

	log.Logger.Info("服务器已经启动🌶~")

	for {
		conn, err := listen.Accept()

		if err != nil {
			log.Logger.Error("accept 发生 error:", err.Error())
		}

		remoteAddrStr := conn.RemoteAddr().String()
		log.Logger.NewUser(remoteAddrStr)

		//保存用户的 ip 和 connect 实例
		ipAndConnMap[remoteAddrStr] = conn

		//开一个协程处理用户连接
		go progress(conn, &ipAndConnMap)
	}
}
