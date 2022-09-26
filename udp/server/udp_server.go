package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	udp, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8888,
	})
	if err != nil {
		logrus.Error("udp", err)
		return
	}
	defer udp.Close()
	for {
		var data [1024]byte
		n, addr, err := udp.ReadFromUDP(data[:])
		if err != nil {
			logrus.Error("read", err)
			continue
		}
		fmt.Printf("data=%v addr =%v  count=%v", string(data[:n]), addr, n)
		_, err = udp.WriteToUDP(data[:n], addr)
		if err != nil {
			logrus.Error("write", err)
			continue
		}

	}
}
