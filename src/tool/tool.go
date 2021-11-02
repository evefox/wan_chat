package tool

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	IntTypeInt8 = iota
	IntTypeInt16
	IntTypeInt32
	IntTypeInt64
)

// IntToBytes 整形转换成字节
func IntToBytes(n int, intType int) []byte {
	x := transInt(n, intType)
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, x)
	if err != nil {
		return nil
	}
	return bytesBuffer.Bytes()
}

func transInt(n int, intType int) interface{} {
	switch intType {
	case IntTypeInt8:
		return int8(n)
	case IntTypeInt16:
		return int16(n)
	case IntTypeInt32:
		return int32(n)
	case IntTypeInt64:
		return int64(n)
	default:
		return int32(n)
	}
}

// BytesToInt 字节转换成整形
func BytesToInt(b []byte) (int, error) {
	bytesBuffer := bytes.NewBuffer(b)
	length := len(b)
	switch length - 1 {
	case IntTypeInt8:
		var x int8
		err := binary.Read(bytesBuffer, binary.BigEndian, &x)
		return int(x), err
	case IntTypeInt16:
		var x int16
		err := binary.Read(bytesBuffer, binary.BigEndian, &x)
		return int(x), err
	case IntTypeInt32:
		var x int32
		err := binary.Read(bytesBuffer, binary.BigEndian, &x)
		return int(x), err
	default:
		return 0, fmt.Errorf("byte lenth not valid")
	}
}
