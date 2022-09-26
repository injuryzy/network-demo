package main

import (
	"bytes"
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"net"
)

func Encode(msg string) ([]byte, error) {
	msgLen := int16(len(msg))
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.LittleEndian, msgLen)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil

}
func main() {
	conn, err := net.Dial("tcp", ":8083")
	if err != nil {
		logrus.Error("dail", err)
	}
	defer conn.Close()
	for i := 0; i < 30; i++ {
		msg := "hello world,hello xiaomotong"
		data, err := Encode(msg)
		if err != nil {
			logrus.Error("encode msg error", err)
			return
		}
		conn.Write(data)
	}

}
