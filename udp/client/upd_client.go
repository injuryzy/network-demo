package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	dial, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8888,
	})
	if err != nil {
		logrus.Error("dial", err)
		return
	}
	defer dial.Close()

	sendData := []byte("hello do something")
	_, err = dial.Write(sendData)
	if err != nil {
		fmt.Println("socket error ", err)
		return
	}
	data := make([]byte, 2048)
	n, addr, err := dial.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read", err)
		return
	}
	fmt.Printf("data == %v  , addr == %v , count == %v\n", string(data[:n]), addr, n)
}
