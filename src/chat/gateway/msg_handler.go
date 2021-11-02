package gateway

import (
	"chat/gen_server"
	"chat/role"
	"fmt"
	pb "github.com/evefox/chat/protos"
)

func HandlerLogin(gwPrt *Gateway, msgID int32, msg interface{}) {
	p := msg.(*pb.Login)
	fmt.Println("HandlerLogin:", msgID, " body:", p)
	if gwPrt.roleID != 0 {
		fmt.Println("login error roleID:", gwPrt.roleID)
		response := pb.LoginResponse{Err: 1}
		sendClient(&response, gwPrt.conn)
		return
	}
	id := int(p.Id)
	server := role.StartRole(id)
	gwPrt.SetRole(id, server)
	response := pb.LoginResponse{Err: 0}
	sendClient(&response, gwPrt.conn)
}

func HandlerLogout(gwPrt *Gateway, msgID int32, msg interface{}) {
	p := msg.(*pb.Logout)
	fmt.Println("HandlerLogout:", msgID, " body:", p)
	response := pb.LogoutResponse{Err: 0}
	sendClient(&response, gwPrt.conn)
}

func HandlerChat(gwPrt *Gateway, msgID int32, msg interface{}) {
	p := msg.(*pb.Chat)
	fmt.Println("HandlerChat:", msgID, " body:", p)
	gen_server.Send(gwPrt.roleServer, p)
}
