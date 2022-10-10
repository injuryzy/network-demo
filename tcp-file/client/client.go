package main

import (
	"github.com/injuryzy/network-demo/common/cerr"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

func main() {
	list := os.Args
	fileName := list[1]

	stat, err := os.Stat(fileName)
	if err != nil {
		logrus.Error(err)
	}
	cerr.Check(err)

	conn, err := net.Dial("tcp", ":8080")
	cerr.Check(err)

	defer conn.Close()

	_, err = conn.Write([]byte(stat.Name()))
	cerr.Check(err)
	var buf = make([]byte, 1024)
	n, err := conn.Read(buf)
	cerr.Check(err)
	if string(buf[:n]) == "ok" {
		sendFile(conn, fileName)
	}
}

func sendFile(conn net.Conn, name string) {
	file, err := os.Open(name)
	cerr.Check(err)
	defer file.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		conn.Write(buf[:n])
	}
}
