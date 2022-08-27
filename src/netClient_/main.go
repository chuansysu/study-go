package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("conn err:", err)
		return
	}
	fmt.Println("conn established:", conn)
	addrRemote := conn.RemoteAddr()
	addrLocal := conn.LocalAddr()
	fmt.Println("remote:", addrRemote.Network(), addrRemote.String())
	fmt.Println("local:", addrLocal.Network(), addrLocal.String())
	for {
		line := Input(conn)
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}
	}
}

func Input(conn net.Conn) string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readstring err:", err)
	}
	line = strings.Trim(line, "\r\n")
	if line == "exit" {
		return line
	}
	n, err := conn.Write([]byte(line + "\n"))
	if err != nil {
		fmt.Println("conn.Write err:", err)
	}
	fmt.Printf("客户端发送了%d字节的数据\n", n)
	return line
}
