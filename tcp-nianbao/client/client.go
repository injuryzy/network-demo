package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8082")
	if err != nil {
		fmt.Println("net dail error", err)
		return
	}
	defer conn.Close()
	fmt.Println("client start ...")
	for i := 0; i < 30; i++ {
		msg := "hello world ,helo xiaomotong "
		conn.Write([]byte(msg))
	}
	fmt.Println("send over")
}
