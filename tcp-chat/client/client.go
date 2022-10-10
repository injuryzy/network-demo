package main

import (
	"bufio"
	"fmt"
	"github.com/injuryzy/network-demo/common/cerr"
	"net"
	"os"
)

func main() {
	dial, err := net.Dial("tcp", ":8080")
	cerr.Check(err)
	defer dial.Close()
	go ReadMessage(dial)
	go InputInfo(dial)
	for {

	}

}

// InputInfo user input info
func InputInfo(dial net.Conn) {
	for {
		line, _, err := bufio.NewReader(os.Stdin).ReadLine()
		cerr.Check(err)
		dial.Write(line)
		if string(line) == "exit" {
			os.Exit(1)
		}
	}

}

func ReadMessage(conn net.Conn) {
	for {
		var buf = make([]byte, 1024)
		n, err := conn.Read(buf)
		cerr.Check(err)
		fmt.Println(string(buf[:n]))
		if string(buf[:n]) == "time out" {
			os.Exit(1)
		}
	}

}
