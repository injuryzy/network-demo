package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net"
)

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(2)
	lengthBuffer := bytes.NewBuffer(lengthByte)
	var length int16
	err := binary.Read(lengthBuffer, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	if int16(reader.Buffered()) < length+2 {
		return "", err
	}
	//read msg
	realData := make([]byte, length+2)
	_, err = reader.Read(realData)
	if err != nil {
		return "", err
	}
	return string(realData[2:]), nil
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg, err := Decode(reader)
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			logrus.Error("read", err)
			return
		}
		fmt.Println("received data: ", msg)
	}

}
func main() {
	listen, err := net.Listen("tcp", ":8083")
	if err != nil {
		logrus.Error("listen ", err)
		return
	}
	defer listen.Close()
	for {
		accept, err := listen.Accept()
		if err != nil {
			logrus.Error("accpet", err)
			continue
		}
		go process(accept)

	}
}
