package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:728")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}
	for {
		if conn, err := listener.Accept(); err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		} else {
			go process(conn)
		}
	}
}
func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}
		fmt.Printf("receive from client, data: %v\n", string(buf[:n]))
		if _, err = conn.Write([]byte("Send From Server")); err != nil {
			fmt.Printf("write to client failed, err: %v\n", err)
			break
		}
	}
}
