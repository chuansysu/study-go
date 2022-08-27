package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	fmt.Println(listen)

	/*
		defer listen.Close()
	*/
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
		} else {
			fmt.Println("conn established:", conn)
			go process(conn)
			//addrRemote := conn.RemoteAddr()
			//addrLocal := conn.LocalAddr()
			//fmt.Println("remote:", addrRemote.Network(), addrRemote.String())
			//fmt.Println("local:", addrLocal.Network(), addrLocal.String())
		}
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		fmt.Println("服务器在等待客户端发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器的read err:", err)
			return
		}
		fmt.Printf("服务器收到长度%d的数据\n", n)
		fmt.Print(string(buf[:n]))
	}
}
