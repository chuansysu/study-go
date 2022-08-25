package main

import (
	"bytes"
	"fmt"
)

func main() {
	/*
		var b = []byte("seafood")
		a := bytes.ToUpper(b)
		fmt.Println(a, b)
		fmt.Println(string(a), string(b))

		c := b[0:4]
		c[0] = 'A'
		fmt.Println(c, b)

		fmt.Println(string(c), string(b))

		bb := []byte("  Hello   World !   ")
		fmt.Printf("%q\n", bytes.Split(bb, []byte{' '}))
		fmt.Printf("%q\n", bytes.Fields(bb))

		f := func(r rune) bool {
			return bytes.ContainsRune([]byte(" !"), r)
		}
		fmt.Printf("%q\n", bytes.TrimFunc(bb, f))
	*/

	/*
		// Reader读取后内容不变，未读取指针移动
		b1 := []byte("Hello World!")
		b2 := []byte("Hello 世界！")
		buf := make([]byte, 6)
		rd := bytes.NewReader(b1)
		rd.Read(buf)
		fmt.Printf("%q\n", buf)
		fmt.Printf("rd:Size:%d, Len:%d\n", rd.Size(), rd.Len()) // Size:12, Len:6
		rd.Read(buf)
		fmt.Printf("%q\n", buf)
		fmt.Printf("rd:Size:%d, Len:%d\n", rd.Size(), rd.Len()) // Size:12, Len:0

		rd.Reset(b2)
		rd.Read(buf)
		fmt.Printf("%q\n", buf)
		fmt.Printf("rd:Size:%d, Len:%d\n", rd.Size(), rd.Len()) // Size:15, Len:9
	*/

	// Buffer读取后内容改变
	rd := bytes.NewBufferString("Hello World!")
	fmt.Printf("rd:%v,%d,%d\n", rd, rd.Len(), rd.Cap()) // len:12,cap:16 //内容向前移动
	buf := make([]byte, 6)
	b := rd.Bytes()
	rd.Read(buf)
	fmt.Printf("rd:%v,%d,%d\n", rd, rd.Len(), rd.Cap()) // len:6,cap:16
	c := rd.Bytes()
	fmt.Printf("%s\n", rd.String())
	fmt.Printf("b:%s,%d,%d\n", b, len(b), cap(b))
	fmt.Printf("c:%s,%d,%d\n", c, len(c), cap(c))
	fmt.Printf("rd:%v,%d,%d\n", rd, rd.Len(), rd.Cap()) // len:6,cap:16
}
