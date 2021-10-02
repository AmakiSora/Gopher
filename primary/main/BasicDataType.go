package main

import (
	"fmt"
	"unsafe"
)

/*
	基本数据类型
		整数型:int,int8,int16,int32,int64,uint,byte
		浮点型:float32,float64
		布尔型(bool):布尔型只有true和false,其他数字等都不行
		字符串(string):在go里字符串算基本数据类型
	注意:go语言没有字符型,使用byte来保存单个字符
*/
func main() {
	/*
		整数型		符号		占用空间	范围(64位系统)
		int(默认)	有符号	8字节	-9223372036854775808 ~ 9223372036854775807
		int8		有符号	1字节	-128 ~ 127
		int16		有符号	2字节	-32768 ~ 32767
		int32		有符号	4字节	-2147483648 ~ 2147483647
		int64		有符号	8字节	-9223372036854775808 ~ 9223372036854775807
		uint		无符号	8字节	0 ~ 18446744073709551615
		uint8		无符号	1字节	0 ~ 255
		uint16		无符号	2字节	0 ~ 65535
		uint32		无符号	4字节	0 ~ 4294967295
		uint64		无符号	8字节	0 ~ 18446744073709551615
		rune		有符号	4字节	-2147483648 ~ 2147483647	等价int32,表示一个Unicode码
		byte		无符号	1字节	0 ~ 255						本质上是uint8,存储字符使用
	*/
	var i int8 //默认为0
	fmt.Println("i的值:", i)
	fmt.Printf("i的类型:%T\n", i)                  //查看数据类型
	fmt.Printf("i的占用空间:%d\n", unsafe.Sizeof(i)) //查看占用空间
	/*
		浮点型		占用空间		范围(64位系统)
		float32		4字节		-3.403E38 ~ 3.403E38
		float64(默认)8字节		-1.798E308 ~ 1.798E308
	*/
	var f = .123
	fmt.Println("f的值:", f)
	fmt.Printf("f的类型:%T\n", f)
	fmt.Printf("f的占用空间:%d\n", unsafe.Sizeof(f)) //查看占用空间
	fmt.Println("1.11e3 ->", 1.11e3)            //科学计数法,等价于1.11 * 10的3次方
	/*
		布尔类型(占1字节)
		Go语言里只允许取值true和false,其他数值(0和非0)不允许
	*/
	var b = true
	fmt.Println("b的值:", b)
	fmt.Printf("b的类型:%T\n", b)
	/*
		字符串
		Go的单个字符一般由byte来存储
		Go的字符串是有字节组成的(不是字符)
		Go语言的字符使用UTF-8编码
	*/
	var c byte = 'a'
	fmt.Println("c的值:", c)
	fmt.Printf("c的类型:%T\n", c)
	str := "字符串"
	fmt.Println("str的值:", str)
	fmt.Printf("str的类型:%T\n", str)

}
