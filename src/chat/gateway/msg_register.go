package gateway

import (
	pb "github.com/evefox/chat/protos"
)

const (
	Login = iota
	LoginResponse
	Logout
	LogoutResponse
	Chat
	ChatResponse
)

func RegisterMsg() {
	//register
	RegisterMessage(Login, &pb.Login{}, HandlerLogin)
	RegisterMessageResponse("LoginResponse", LoginResponse)
	RegisterMessage(Logout, &pb.Logout{}, HandlerLogout)
	RegisterMessageResponse("LogoutResponse", LogoutResponse)
	RegisterMessage(Chat, &pb.Chat{}, HandlerChat)
	RegisterMessageResponse("ChatResponse", ChatResponse)
}
