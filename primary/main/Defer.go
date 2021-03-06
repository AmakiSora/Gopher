package main

import "fmt"

/*
	defer关键字
	当函数执行到最后时,这些defer语句会按照逆序执行,最后该函数返回(有点类似java的try-catch中的finally)
	使用场景:当你想关闭资源的时候,随手写defer,因为defer的机制会在执行完毕后关闭资源,省事

	也可以结合recover来捕获错误,进行异常处理
*/

func main() {
	fmt.Println("结果:", subtraction(10, 5))
	deferRecover()
}
func subtraction(n1 int, n2 int) (sub int) {
	//在go语言中,程序遇到defer关键字,不会立即执行defer后的语句,而是将defer后的语句压入一个栈中,然后继续执行函数后面的语句
	defer fmt.Println("n1 =", n1) //会把变量的值也压入栈中,最后会输出n1 = 10,而不是n1 = 20
	defer fmt.Println("n2 =", n2)
	n1 += 10
	n2 += 5
	sub = n1 - n2
	fmt.Println("n1 - n2 =", sub)
	//在函数执行完毕后,从栈中取出语句,按照先进后出的规则执行语句
	return
}

func deferRecover() {
	//利用defer+recover来捕获错误
	defer func() {
		//调用recover内置函数,可以捕获错误
		err := recover()
		//如果没有捕获错误,返回值为nil
		if err != nil {
			fmt.Println("出错了!错误为:", err)
		}
	}()
	a := 1
	b := 0
	fmt.Println(a / b)
}
