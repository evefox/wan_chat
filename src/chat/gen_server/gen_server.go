package gen_server

import "fmt"

const (
	call = iota
	send
)

const (
	Reply = iota
	NoReply
)

type GenServerMod interface {
	Init(args interface{}) (ok bool)
	DoHandleCall(msg interface{}) (result GenCallResult)
	DoHandleInfo(msg interface{}) (ok bool)
	Terminate()
}

type GenServerReply struct {
	result int
	reply  interface{}
}

type GenCallResult struct {
	ReplyType int
	Result    int
	Reply     interface{}
}

type GenServerMsg struct {
	msgType   int
	msg       interface{}
	replyChan chan GenServerReply
}

func Call(caller chan GenServerMsg, msg interface{}) (result int, reply interface{}) {
	replyChan := make(chan GenServerReply, 0)
	caller <- GenServerMsg{msgType: call, msg: msg, replyChan: replyChan}
	for {
		select {
		case replyMsg := <-replyChan:
			return replyMsg.result, replyMsg.reply
		}
	}
}

func Send(caller chan GenServerMsg, msg interface{}) {
	caller <- GenServerMsg{msgType: send, msg: msg, replyChan: nil}
}

func Start(mod GenServerMod, recv chan GenServerMsg, args interface{}) {
	go startNewServer(mod, recv, args)
}

func startNewServer(mod GenServerMod, recv chan GenServerMsg, args interface{}) {
	if ok := mod.Init(args); ok {
		enterLoop(mod, recv)
	}
}

func enterLoop(mod GenServerMod, recv chan GenServerMsg) {
	for {
		select {
		case msg := <-recv:
			if ok := doMsg(msg, mod); !ok {
				fmt.Printf("stop\n")
				mod.Terminate()
				return
			}
		}
	}
}

func doMsg(msg GenServerMsg, mod GenServerMod) (result bool) {
	if msg.msgType == call {
		callResult := mod.DoHandleCall(msg)
		if callResult.ReplyType == Reply {
			msg.replyChan <- GenServerReply{result: callResult.Result, reply: callResult.Reply}
			return true
		} else if callResult.ReplyType == NoReply {
			return true
		} else {
			return false
		}
	} else {
		return mod.DoHandleInfo(msg)
	}
}
