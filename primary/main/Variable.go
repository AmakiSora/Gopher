package main

import "fmt"

/*
	变量
*/
func main() {
	//变量的声明(int默认值为0)
	var a int

	//go会根据值自行判断变量类型
	var b = 2
	var c = "233"

	//:=符号的左边不能是已经声明过的
	d := 1.1

	//声明多个变量
	var e1, e2, e3 int8
	f1, f2, f3 := 5, 6, true
	var (
		g1 = 8
		g2 = 9
		g3 = ""
	)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e1, e2, e3)
	fmt.Println("f = ", f1, f2, f3)
	fmt.Println("g = ", g1, g2, g3)

}
