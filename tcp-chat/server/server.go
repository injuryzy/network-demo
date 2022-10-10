package main

import (
	"github.com/injuryzy/network-demo/common/cerr"
	user "github.com/injuryzy/network-demo/tcp-chat"
	"net"
	"sync"
	"time"
)

var OnlionMap = make(map[string]user.User)

var close = make(chan int)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	defer listen.Close()
	cerr.Check(err)
	for {
		conn, err := listen.Accept()
		if err != nil {
			cerr.Check(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	group := sync.WaitGroup{}
	group.Add(1)
	conn.Write([]byte("请输入用户名:"))
	defer conn.Close()
	name := ReadMessage(conn)
	user := user.User{
		Name: name,
		Addr: conn.RemoteAddr().String(),
		C:    make(chan string),
	}
	OnlionMap[name] = user
	conn.Write([]byte("欢迎" + name + "进入聊天室"))
	go ChanleMsg(conn, user)
	go WriteMsgToClient(group, &user, conn)

	group.Wait()
}

func ChanleMsg(conn net.Conn, u user.User) {
	for {
		msg := ReadMessage(conn)
		if msg == "exit" || <-close == 1 {
			delete(OnlionMap, u.Name)
			conn.Close()
			break
		}
		if msg == "online" {
			for _, v := range OnlionMap {
				conn.Write([]byte(v.Name + ":" + v.Addr))
			}
			continue
		}

		for _, v := range OnlionMap {
			v.C <- u.Name + ":" + msg
		}
	}
}

func WriteMsgToClient(group sync.WaitGroup, u *user.User, conn net.Conn) {
	group.Done()
	for {
		select {
		case msg := <-u.C:
			conn.Write([]byte(msg))
		case <-time.After(10 * time.Second):
			conn.Write([]byte("time out"))
			close <- 1
		}
	}
}

//ReadMessage  conn read message
func ReadMessage(conn net.Conn) string {
	var buf = make([]byte, 1024)
	n, err := conn.Read(buf)
	cerr.Check(err)
	return string(buf[:n])
}
