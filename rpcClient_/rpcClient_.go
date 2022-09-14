package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	cli, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	//var request = "helloworld"
	var reply string
	err = cli.Call("HelloSerer.Hello", "helloworld", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
