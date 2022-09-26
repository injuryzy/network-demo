package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"strings"
)

//与服务端创建链接
//读写数据
//关闭

func main() {
	dial, err := net.Dial("tcp", ":8081")
	if err != nil {
		logrus.Error("dial", err)
		return
	}
	defer dial.Close()
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err = dial.Write([]byte(inputInfo))
		if err != nil {
			logrus.Error(err)
			return
		}
		var buf [512]byte
		n, err := dial.Read(buf[:])
		if err != nil {
			logrus.Error(err)
		}
		fmt.Println(string(buf[:n]))
	}
}
