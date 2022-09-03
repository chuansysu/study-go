package main

import (
	"fmt"
	"unsafe"
)

type s struct{ i int }

const VERSION = "1.0"

func main() {
	//log_.DemoLog()
	//yaml_.DemoYaml()
	//filepath_.DemoFilePath()
	//io_.DemoIO()
	// q := [3]int{1, 2, 3}[1:]
	fmt.Printf("%p\n", &s{})
	fmt.Println(unsafe.Pointer(&s{}))
	q := &s{}
	fmt.Printf("%p,%v\n", q, q)
	fmt.Printf("%p\n", &[]int{1, 2, 3})
}
