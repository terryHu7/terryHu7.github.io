package server

import (
	"bufio"
	"fmt"
	"net"
)

func TcpServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:728")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close() //TIME_WAIT
	for {
		reader := bufio.NewReader(conn)
		var buf [256]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("client: ", recvStr)
		conn.Write([]byte(recvStr))
	}
}

//https://www.liwenzhou.com/posts/Go/15_socket/
