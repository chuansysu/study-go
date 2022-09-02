package io_

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

// Reader 在输入流结束时会返回一个非零的字节数，同时返回的 err 不是 EOF 就是 nil。
// 无论如何，下一个 Read 都应当返回 0, EOF。

// 调用者应当总在考虑到错误 err 前处理 n > 0 的字节。
// 这样做可以在读取一些字节，以及允许的 EOF 行为后正确地处理 I/O 错误。

func demoReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func demoWriteTo(p []byte) (int, error) {
	var writer io.Writer = os.Stdout
	n, err := writer.Write(p)
	return n, err
}

// ReadAt 的客户端可对相同的输入源并行执行 ReadAt 调用。
// ReaderAt 接口使得可以从指定偏移量处开始读取数据。

func demoReadAt() {
	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s,%d\n", p, n)
	fmt.Println(reader.Len())

	p2 := make([]byte, 6)
	n2, err := reader.ReadAt(p2, 5)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s,%d\n", p2, n2)
}

// 若区域没有重叠，WriteAt 的客户端可对相同的目标并行执行 WriteAt 调用。

func demoWriteAt() {
	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余")
	n, err := file.WriteAt([]byte("Go语言中文网"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

// 注意：ReadFrom 方法不会返回 err == EOF。

// 如果不通过 ReadFrom 接口来做这件事，而是使用 io.Reader 接口，我们有两种思路：

// 1.先获取文件的大小（File 的 Stat 方法），之后定义一个该大小的 []byte，通过 Read 一次性读取
// 2.定义一个小的 []byte，不断的调用 Read 方法直到遇到 EOF，将所有读取到的 []byte 连接到一起

// 提示
// 通过查看 bufio.Writer 或 strings.Buffer 类型的 ReadFrom 方法实现，
// 会发现，其实它们的实现和上面说的第 2 种思路类似。

func demoReaderFrom() {
	file, err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}

func demoWriterTo() {
	reader := bytes.NewReader([]byte("Go语言中文网\n"))
	reader.WriteTo(os.Stdout) // bytes.Reader实现WriteTo接口
}

func demoTeeReader() {
	/*
		out, err := os.Create("test.tmp")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		rsp, err := http.Get("https://down.qq.com/qqweb/PCQQ/PCQQ_EXE/PCQQ2020.exe")
		if err != nil {
			log.Fatal(err)
		}
		defer rsp.Body.Close()
		teeReader := &Speeder{}
		go teeReader.Show()
		io.Copy(out, io.TeeReader(rsp.Body, teeReader))
	*/

	reader := io.TeeReader(strings.NewReader("Go语言中文网"), os.Stdout)
	p := make([]byte, 20)
	reader.Read(p)
	fmt.Println("str:", p)
}

type Speeder struct {
	count int64
}

func (s *Speeder) Write(b []byte) (int, error) {
	c := len(b)
	atomic.AddInt64(&s.count, int64(c))
	return c, nil
}

func (s *Speeder) Show() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		fmt.Printf("\r%.2fkb/s", float64(atomic.LoadInt64(&s.count))/float64(1024))
		atomic.StoreInt64(&s.count, 0)
	}
}

func DemoIO() {
	//p, err := io_.ReadFrom(os.Stdin, 10)
	//fmt.Println(p, err)

	n, err := demoWriteTo([]byte("hello world!"))
	fmt.Println(n, err)

	demoReadAt()
	demoWriteAt()
	demoReaderFrom()
	fmt.Println()
	demoWriterTo()
	demoTeeReader()
}
