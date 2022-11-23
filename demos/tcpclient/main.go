package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("�˿���(www.haicoder.net)")
	conn, err := net.Dial("tcp", "localhost:728")
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err: %v\n", err)
			break
		}
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}
		//��������
		if _, err = conn.Write([]byte(trimmedInput)); err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
		//��������
		var recvData = make([]byte, 1024)
		if _, err = conn.Read(recvData); err != nil {
			fmt.Printf("Read failed , err : %v\n", err)
			break
		}
		fmt.Printf("Receive data : %v\n", string(recvData))
	}
}
