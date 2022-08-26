// bufio通过缓存来提高效率，缓存放在主存中。
// Golang的bufio包实现了带缓存的I/O读写操作，用来帮助处理I/O缓存。
// 通过缓存可以提高效率，把文件读取进缓存（内存）后再读取的时候可避免文件系统的I/O，从而提高速度。
// 同理，当进行写操作时会先把文件写入缓存（内存），然后由缓存写入文件系统。
// 缓冲区的设计目的是为了存储多次的写入，最后一次性将缓冲区内容写入文件。

package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	b, _ := br.Peek(5)
	fmt.Printf("%s\n", b)

	b[0] = 'a'
	//b, _ = br.Peek(5)
	fmt.Printf("b:%s\n", b)
	fmt.Printf("s:%d\n", s.Size())
	fmt.Printf("br:%d\n", br.Size())
}
