package endian_

import (
	"fmt"
	"unsafe"
)

// BigEndian:指数据的高字节保存在内存的低地址中，而数据的低字节保存在内存的高地址中
// LittleEndian:指数据的高字节保存在内存的高地址中，而数据的低字节保存在内存的低地址中

// 比如对于数据:0x12345678，从高字节到低字节为：12 34 56 78，从低字节到高字节为:78 56 34 12。
// 按照大端模式从低位 buf[0]到高位 buf[3]则应该为：12， 34， 56， 78。
// 按照小端模式从低位 buf[0]到高位 buf[3]则应该为: 78，56，34，12。

func DemoEndian() {
	var i int = 0x12345678                  //定义数据
	const size int = int(unsafe.Sizeof(i))  //获取i的长度,
	ps := (*[size]byte)(unsafe.Pointer(&i)) //此处size必须为const类型，不然会报错
	fmt.Printf("%T\n", ps)                  //*[8]byte类型
	fmt.Println(*ps)                        //打印值 //打印存储地址
	fmt.Println(&ps[0])
	fmt.Println(&ps[1])
	fmt.Println(&ps[2])
	fmt.Println(&ps[3])
	if ps[0] == 0x78 {
		//小端模式则ps[0]低位存放的是低字节0x78,十进制则为120，  满足低位存放低字节，    存储为  0x78563412不利用阅读，但方便计算机进行运算。
		fmt.Println("系统为小端模式")
	} else {
		//大端模式则ps[0]低位存放的是高字节12,十进制为18,满足低位存放高字节，存储为ox12345678，方便阅读，但不方便计算机进行运算
		fmt.Println("系统为大端模式")
	}
}
