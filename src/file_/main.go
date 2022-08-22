package main

import (
	"fmt"
	"io"
	"os"
)

func foo() {
	file, err := os.Open("./test")
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			fmt.Printf("%T,%v\n", err, err)
		}
		fmt.Println("open file error", err)
		return
	}
	defer file.Close()

	fmt.Printf("%T,%+#[1]v\n", file)

	var tmp = make([]byte, 128)
	fmt.Printf("%v,%d,%d\n", tmp, len(tmp), cap(tmp))
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Printf("%s\n", string(tmp))
	/*
		reader := bufio.NewReader(file)
		for {
			str, err := reader.ReadString('\n')
			fmt.Print(str)
			if err == io.EOF {
				break
			}
		}
		fmt.Println()
		fmt.Println("文件读取结束")
	*/
	//content, _ := ioutil.ReadFile("./test")
	//fmt.Print(string(content))
}

func main() {
	foo()
}
