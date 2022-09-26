package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/injuryzy/go-tool/rlog"
	"github.com/injuryzy/go-tool/robot"
	"github.com/sirupsen/logrus"
)

// 监听端口
// 接受客户端连接
// 处理链接
// 关闭
var log *logrus.Logger

func init() {
	log = rlog.NewLog()
	newRobot := robot.NewRobot("https://open.feishu.cn/open-apis/bot/v2/hook/136aee8f-4d5a-4ea5-82a4-e6cdb1295224", "")
	log.AddHook(rlog.NewRobotHook(&newRobot, logrus.ErrorLevel))
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Error("链接失败", err)
	}

	for {
		conn, err2 := listen.Accept()
		if err2 != nil {
			log.Error("等待失败", err2)
		}
		go Porcess(conn)
	}
}

func Porcess(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [256]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			log.Error("读取数据失败", err)
		}
		recvDate := buf[:n]
		fmt.Println("receDate", string(recvDate))
		conn.Write(recvDate)
	}
}
