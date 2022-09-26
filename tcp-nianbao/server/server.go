package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net"
)

func main() {
	tcp, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8082,
	})
	if err != nil {
		logrus.Error("listen error ", err)
	}
	for {
		accept, err := tcp.Accept()
		if err != nil {
			logrus.Error("accpet error ", err)
			continue
		}
		go process(accept)
	}

}

// 处理进程
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [2048]byte
	for {
		n, err := reader.Read(buf[:])
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			logrus.Error("read ", err)
			break
		}
		revcStr := string(buf[:n])
		fmt.Println("receive data :%s \n\n", revcStr)
	}
}
