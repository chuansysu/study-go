package main

import "fmt"

/*
type Reader interface {
	Read([]byte) (n int, err error)
}

type Writer interface {
	Write([]byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

// whence
const (
	SeekStart   = 0
	SeekCurrent = 1
	SeekEnd     = 2
)
*/

func f(b []byte) {
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &b)
}
func main() {
	s1 := []byte{1}
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", &s1)
	f(s1)

}
