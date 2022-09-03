package log_

import (
	"log"
	"os"
)

func DemoLog() {
	//log.Println("test")
	//log.Printf("hello:%d\n", 10)
	//log.Printf("33333%s:555555\n", "hello")
	//log.Fatalln("kkkkkk")
	//log.Fatal("panic...")
	//log.Panicln("this is panic test.")

	logfile, err := os.OpenFile("my.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("fail to create log file!")
	}
	defer logfile.Close()

	l := log.New(logfile, "", log.LstdFlags)
	l.Println("test")
	num := 5
	l.Println("test %d", num)
}
