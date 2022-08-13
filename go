
GOROOT:/usr/local/go
GOPATH:/home/gchuan/go

1: GO111MODULE=off

目录：
	/home/gchuan/go/bin
	/home/gchuan/go/pkg
	/home/gchuan/go/src
项目：
	/home/gchuan/go/src
		|
	   myproject
		  |
		 main.go
		 utils
		    |
		   utils.go
文件:
	myproject/utils/utils.go:
		package utils
		import "fmt"
		func Hello(){
			fmt.Println("msg in utils.")
		}
	myproject/main.go:
		package main
		import(
			"fmt"
			"myproject/utils"
		}
		func main(){
			fmt.Println("Hello World!")
			utils.Hello()
		}
编译：
***网络上很多资料是错误的，不知道是不是版本原因***
	1) 在myproject/目录执行go run main.go
	2) 在任意目录执行go run myproject都正常输出
	3) 在myproject/目录执行go build main.go则在/myproject/目录生成main(***注意名字***),无其他生成
	4) 在myproject/目录执行go build则在/myproject/目录生成myproject(***注意名字***),无其他生成
	5) 在任意目录执行go build myproject则在/执行命令的当前目录生成myproject(***注意名字***),无其他生成
	6) 任何目录执行go install myproject都可以在$GOPATH/bin/目录下生成main，pkg目录无输出
	7) 任何目录执行go install myproject/utils都可以在$GOPATH/pkg/目录下生成pkg/linux_amd64/myproject/utils.a
	总结:某些命令编译器会自动到$GOPATH/src/目录下面查找项目名称，比如go install myproject















	
