package gateway

import (
	"chat/gen_server"
	pb "chat/protos"
	"chat/role"
	"fmt"
)

func HandlerLogin(gwPrt *Gateway, msgID int32, msg interface{}) {
	if gwPrt.roleID != 0 {
		response := pb.LoginResponse{Err: 1}
		sendClient(&response, gwPrt.conn)
		return
	}
	p := msg.(*pb.Login)
	id := int(p.Id)
	server := role.StartRole(id)
	gwPrt.SetRole(id, server)
	fmt.Println("message handler msgid:", msgID, " body:", p)
	response := pb.LoginResponse{Err: 0}
	sendClient(&response, gwPrt.conn)
}

func HandlerLogout(gwPrt *Gateway, msgID int32, msg interface{}) {
	p := msg.(*pb.Logout)
	fmt.Println("message handler msgid:", msgID, " body:", p)
	response := pb.LogoutResponse{Err: 0}
	sendClient(&response, gwPrt.conn)
}

func HandlerChat(gwPrt *Gateway, msgID int32, msg interface{}) {
	p := msg.(*pb.Chat)
	gen_server.Send(gwPrt.roleServer, p)
}
