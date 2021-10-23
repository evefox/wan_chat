package gateway

import (
	pb "chat/protos"
	"fmt"
)

func HandlerLogin(msgID int32, msg interface{}) {
	p := msg.(*pb.Login)
	fmt.Println("message handler msgid:", msgID, " body:", p)
}

func HandlerLogout(msgID int32, msg interface{}) {
	p := msg.(*pb.Logout)
	fmt.Println("message handler msgid:", msgID, " body:", p)
}
