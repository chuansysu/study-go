package main

import (
	"fmt"
	"log"
	"myprotopb"

	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("test")
	login := &myprotopb.Login{
		Account: "hellokitty",
		UserId:  12345,
	}
	fmt.Println(login)

	pb, err := proto.Marshal(login)
	if err != nil {
		log.Fatal("proto err:", err)
	}
	fmt.Println(pb)
}
