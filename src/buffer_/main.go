package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("hello"))
	fmt.Println(b, b.Len(), b.Cap())
	fmt.Println(b.String())

	b1 := new(bytes.Buffer)
	fmt.Println(b1)
	b1.Write([]byte("hellokitty"))
	fmt.Println(b1.Len(), b1.Cap())
	fmt.Println(b1)

	bytesBUffer := bytes.NewBuffer([]byte{})
	fmt.Println(bytesBUffer.Cap())
}
