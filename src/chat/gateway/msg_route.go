package gateway

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"reflect"
)

type MessageHandler func(msgID int32, msg interface{})

type MessageInfo struct {
	msgType    reflect.Type
	msgHandler MessageHandler
}

var (
	msgMap         = make(map[int32]MessageInfo)
	msgResponseMap = make(map[string]int32)
)

func RegisterMessage(msgID int32, msg interface{}, handler MessageHandler) {
	var info MessageInfo
	info.msgType = reflect.TypeOf(msg.(proto.Message))
	info.msgHandler = handler

	msgMap[msgID] = info
}

func RegisterMessageResponse(msgName string, msgID int32) {
	msgResponseMap[msgName] = msgID
}

func MsgRoute(msgID int32, data []byte) error {
	if info, ok := msgMap[msgID]; ok {
		msg := reflect.New(info.msgType.Elem()).Interface()
		err := proto.Unmarshal(data, msg.(proto.Message))
		if err != nil {
			return err
		}
		info.msgHandler(msgID, msg)
		return err
	}
	return errors.New("not found msgID")
}
