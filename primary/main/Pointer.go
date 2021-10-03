package main

import "fmt"

/*
	指针
*/
func main() {
	var i int = 233
	fmt.Println("i的值:", i)
	fmt.Println("i的地址:", &i)       //&为取地址符
	var pi *int = &i               //指针,指向i
	fmt.Printf("pi的值:%v\n", pi)    //指针的值
	fmt.Printf("pi的地址:%v\n", &pi)  //指针本身也有地址
	fmt.Printf("pi指向的值:%v\n", *pi) //取指针指向地址的值
	*pi = 666                      //通过指针修改i的值
	fmt.Printf("i的值:%d \n", i)     //变量i的值
}
