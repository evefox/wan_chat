package main

import (
	"chat/chat_server"
	"chat/gateway"
)

func main() {
	gateway.RegisterMsg()

	go gateway.Listen()
	chat_server.Start()
}
