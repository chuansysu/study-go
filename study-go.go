package main

import (
	"study-go/endian_"

	flag "github.com/spf13/pflag"
)

var ip *string = flag.String("ip", "127.0.0.1", "bind ip address")
var port int
var isAuth bool

func init() {
	flag.IntVar(&port, "port", 8080, "listening port")
	flag.BoolVarP(&isAuth, "isAuth", "a", false, "whether auth enable")
}

// study-go.exe version --ip "127.0.0.1" --port 9999 -a=false

func main() {
	flag.Parse()
	//fmt.Println("ipAddr has value ", *ip)
	//fmt.Println("tcpPort has value ", port)
	//fmt.Println("isAuth has value ", isAuth)
	//log_.DemoLog()
	//yaml_.DemoYaml()
	//filepath_.DemoFilePath()
	//io_.DemoIO()
	//viper_.DemoViper()
	//viper := viper_.GetViper()
	//fmt.Println(viper.AllKeys())
	//cobra_.Execute()
	endian_.DemoEndian()
}
