package gateway

import (
	"bytes"
	"chat/gen_server"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
)

const singleMsgSaveLen = 2

type Gateway struct {
	roleID     int
	roleServer chan gen_server.GenServerMsg
	conn       *net.TCPConn
}

// 网关实例
func gatewayInit(conn *net.TCPConn) {
	//声明一个管道用于接收客户端消息的数据
	readerChannel := make(chan []byte, 64)
	sendChannel := make(chan proto.Message, 64)

	//tcp连接的地址
	ipStr := conn.RemoteAddr().String()

	defer func() {
		fmt.Println(" Disconnected : " + ipStr)
		close(readerChannel)
		err := conn.Close()
		if err != nil {
			return
		}
	}()

	go loopReceive(conn, readerChannel)

	gwPrt := new(Gateway)
	gwPrt.conn = conn
	for {
		select {
		case data := <-readerChannel:
			msgData := data[4:]
			err := MsgRoute(gwPrt, int32(bytesToInt(data[:4])), msgData)
			if err != nil {
				return
			}
		case msg := <-sendChannel:
			err := sendClient(msg, conn)
			if err != nil {
				return
			}
		}
	}

}

// 接收客户端发来的消息并分包
func loopReceive(conn *net.TCPConn, readerChannel chan []byte) {
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)
	buffer := make([]byte, 1024)
	//接收并返回消息
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		tmpBuffer = unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}

//分包
func unpack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+singleMsgSaveLen {
			break
		}
		messageLength := bytesToInt(buffer[i : i+singleMsgSaveLen])
		if length < i+singleMsgSaveLen+messageLength {
			break
		}
		data := buffer[i+singleMsgSaveLen : i+singleMsgSaveLen+messageLength]
		readerChannel <- data

		i += singleMsgSaveLen + messageLength - 1
	}

	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}

// 整形转换成字节
func intToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, x)
	if err != nil {
		return nil
	}
	return bytesBuffer.Bytes()
}

// 字节转换成整形
func bytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	err := binary.Read(bytesBuffer, binary.BigEndian, &x)
	if err != nil {
		return 0
	}

	return int(x)
}

// 发送消息到客户端
func sendClient(msg proto.Message, conn *net.TCPConn) error {
	msgBin, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), " send client error: ", err)
		return err
	}
	reflect := proto.MessageReflect(msg)
	msgName := reflect.Descriptor().FullName()
	if msgID, ok := msgResponseMap[string(msgName)]; ok {
		length := len(msgBin)
		msgIDBin := intToBytes(int(msgID))
		tmpBuffer := append(intToBytes(length), msgIDBin[:]...)
		scBin := append(tmpBuffer, msgBin[:]...)
		_, err := conn.Write(scBin)
		if err != nil {
			return err
		}
		return err
	}

	return errors.New("not found msgID")
}

func (g *Gateway) SetRole(id int, server chan gen_server.GenServerMsg) {
	g.roleID = id
	g.roleServer = server
}
