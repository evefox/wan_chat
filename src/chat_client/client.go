package main

import (
	"bufio"
	"fmt"
	pb "github.com/evefox/chat/protos"
	"github.com/evefox/chat/tool"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	Login = iota
	LoginResponse
	Logout
	LogoutResponse
	Chat
	ChatResponse
)

func main() {
	server := "127.0.0.1:15500"
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Printf("Fail error: %s\n", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Printf("Fail error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("connect success")

	login(conn)
}

func login(conn net.Conn) {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("enter id:")
	idStr, err := inputReader.ReadString('\n')
	idStr = strings.TrimSuffix(idStr, "\n")
	if err != nil {
		fmt.Printf("idStr error: %s", err.Error())
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("id error: %s", err.Error())
		return
	}
	fmt.Println("enter name:")
	name, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("name error: %s", err.Error())
		return
	}
	login := pb.Login{
		Id:   int32(id),
		Name: name,
	}
	err = sendServer(conn, Login, &login)
	if err != nil {
		return
	}
	waitLogin(conn)
	sender(conn)
}

func waitLogin(conn net.Conn) {

}

func sender(conn net.Conn) {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("enter chat:")
		msg, err := inputReader.ReadString('\n')
		if err != nil {
			return
		}
		chat := pb.Chat{Content: msg}
		err = sendServer(conn, Chat, &chat)
		if err != nil {
			return
		}
	}
}

func sendServer(conn net.Conn, msgID int, msg proto.Message) (err error) {
	msgBin, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	length := len(msgBin)
	msgIDBin := tool.IntToBytes(msgID, tool.IntTypeInt16)
	tmpBuffer := append(tool.IntToBytes(length+4, tool.IntTypeInt16), msgIDBin[:]...)
	csBin := append(tmpBuffer, msgBin[:]...)
	fmt.Printf("%v\n", csBin)
	_, err = conn.Write(csBin)
	return
}
