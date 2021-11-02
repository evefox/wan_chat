package role

import (
	"chat/gen_server"
	"fmt"
	"reflect"
)

type Role struct {
	RoleID int
}

func StartRole(id int) (roleChan chan gen_server.GenServerMsg) {
	r := Role{RoleID: id}
	recv := make(chan gen_server.GenServerMsg, 100)
	gen_server.Start(r, recv, nil)
	return recv
}

func (r Role) Init(_ interface{}) (ok bool) {
	fmt.Printf("start role:%d\n", r.RoleID)
	return true
}

func (r Role) DoHandleCall(msg interface{}) (result gen_server.GenCallResult) {
	fmt.Printf("%d receive call:%v\n", r.RoleID, msg)
	return gen_server.GenCallResult{ReplyType: gen_server.Reply, Result: 1, Reply: "get call"}
}

func (r Role) DoHandleInfo(msg interface{}) (ok bool) {
	fmt.Println("receive msg:", r.RoleID, ",", msg, ",", reflect.TypeOf(msg))
	return true
}

func (r Role) Terminate() {
	fmt.Printf("%d stop\n", r.RoleID)
}
