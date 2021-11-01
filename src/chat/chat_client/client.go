package main

import (
	"bufio"
	pb "chat/protos"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
)

func main() {
	server := "127.0.0.1:12500"
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
	id, err := inputReader.ReadString('\n')
	if err != nil {
		return
	}
	fmt.Println("enter name:")
	name, err := inputReader.ReadString('\n')
	if err != nil {
		return
	}
	login := pb.Login{
		Id:   int32(id), // todo
		Name: name,
	}
	data, err := proto.Marshal(&login)
	if err != nil {
		return
	}
	_, err = conn.Write(data)
	if err != nil {
		return
	}
	sender(conn)
}

func sender(conn net.Conn) {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("enter chat:")
		msg, err := inputReader.ReadString('\n')
		if err != nil {
			return
		}
		chat := pb.Chat{
			Content: msg,
		}
		data, err := proto.Marshal(&chat)
		_, err = conn.Write(data)
		if err != nil {
			return
		}
	}
}
