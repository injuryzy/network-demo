package main

import (
	"github.com/injuryzy/network-demo/common/cerr"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

func main() {

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Error(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			logrus.Error(err)
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	var buf = make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		logrus.Error(err)
	}
	_, err = os.Create(string(buf[:n]))
	cerr.Check(err)
	conn.Write([]byte("ok"))
	RecvFile(conn, string(buf[:n]))

}

func RecvFile(conn net.Conn, s string) {
	file, err := os.OpenFile(s, os.O_WRONLY|os.O_CREATE, 0666)
	cerr.Check(err)
	defer file.Close()
	var buf = make([]byte, 1024*4)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		file.Write(buf[:n])
	}
}
