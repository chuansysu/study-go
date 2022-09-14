package main

import (
	"fmt"

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

type T interface {
	String() string
}

type S struct {
	s string
}

func (s *S) String() string {
	return s.s
}

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
	//endian_.DemoEndian()

	var s []int = []int{1, 2, 3}
	fmt.Println(s, len(s), cap(s))
	s = append(s, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println(s, len(s), cap(s))

	/*

		var k *S
		fmt.Println(k)
		var i T = k
		t, ok := i.(*S)
		if ok {
			fmt.Println("t", t)
			if t != nil {
				fmt.Println(t.s)
			}
		}
		fmt.Printf("%T,%v\n", i, i)
		p := unsafe.Pointer(&i)
		fmt.Println(p)

		var ni uint = 33
		sqlQuote(ni)

		type tint int
		var ti tint
		var ki int = 3
		ti = tint(ki)
		fmt.Println(ti)

		var a = &A{i: 10}
		var b *B = (*B)(unsafe.Pointer(a))
		fmt.Println(b)
		fmt.Printf("%T,%v\n", a, a)
		fmt.Printf("%T,%v\n", b, b)

	*/

	//var pa *A = &a
	//fmt.Println(pa)
	//fmt.Printf("%T,%p\n", pa, pa)
	//var pb *B = (*B)(unsafe.Pointer(pa))
	//fmt.Printf("%T,%p\n", pb, pb)

	/*
		var ss string
		ch := make(chan int, 10)
		var arr [4]int
		var m map[string]int
		var sl []byte
		var sa A
		var pi *int
		type pp *int
		var ppt pp
		fmt.Println(unsafe.Sizeof(ppt))
		fmt.Println(unsafe.Sizeof(ss))                  // 16
		fmt.Println(unsafe.Sizeof(unsafe.Pointer(&ss))) // 8
		fmt.Println(unsafe.Sizeof(pi))                  // 8
		fmt.Println(unsafe.Sizeof(sa))                  // 4
		fmt.Println(unsafe.Sizeof(sl))                  // 24
		fmt.Println(unsafe.Sizeof(m))                   // 8	// pointer
		fmt.Println(unsafe.Sizeof(ch))                  // 8	// pointer
		fmt.Println(unsafe.Sizeof(arr))                 // 32
		fmt.Println(len(arr), cap(arr))                 // 4,4
	*/
}

type A struct {
	i int32
}

type B struct {
	a int
	b int
}

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		fmt.Printf("%T,%v\n", x, x)
		return fmt.Sprintf("%d", x)
	case string:
		fmt.Println(x)
		return fmt.Sprintf("%s", x)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
