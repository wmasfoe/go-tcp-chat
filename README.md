## 基于 TCP 协议的在线聊天室

使用 GoLang 实现的一个小demo🙈

客户端启动，也可以启动多个：

```sh
cd chat-client
go run .
```

服务端启动：

```sh
cd chat-server
go run .
```

食用指南

- 默认在全局喊话，所有连接的客户端都可看到消息。
- 可以使用 `->ip:port:message` 指令，向指定的 ip 地址发送私聊，这条消息在全局不可见。
- 输入 `:exit` 指令可以退出聊天。
