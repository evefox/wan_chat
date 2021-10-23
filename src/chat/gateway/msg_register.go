package gateway

import (
	pb "chat/protos"
)

const (
	Login = iota
	LoginResponse
	Logout
	LogoutResponse
)

func RegisterMsg() {
	//register
	RegisterMessage(Login, &pb.Login{}, HandlerLogin)
	RegisterMessageResponse("LoginResponse", LoginResponse)
	RegisterMessage(Logout, &pb.Logout{}, HandlerLogout)
	RegisterMessageResponse("LogoutResponse", LogoutResponse)
}
