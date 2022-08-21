package main

import (
	"fmt"
)

type Man struct {
	name string
}

// 输出 My name is:SNS
func (m *Man) String() string {
	return "My name is:" + m.name
}

func f(a fmt.Stringer) {
	fmt.Println(a.String())
}

// 输出{SNS}
/*
func (m *Man) String() string {
	return "My name is:" + m.name
}*/

func main() {
	var m Man
	m.name = "SNS"
	fmt.Println(m)
	fmt.Printf("%s\n", m)
	//f(m) //error:没有对应方法func(m Man)String()string
}
