package main

import (
	"fmt"
	"strconv"
)

/*
	函数
	func 函数名(形参列表)(返回值类型列表){return 返回值列表}
	函数名首字母大写该函数可以被本包文件和其他包文件使用(类似java的public)
	函数名首字母小写该函数只能被本包使用,其他包不能使用(类似java的private)
	go语言不支持重载
	go语言里,函数也是一种数据类型,可以赋值给一个变量,赋值后的变量就是一个函数类型的变量了
*/
//返回值只有一个的时候可以不写括号,没有返回值可以不写
func addition(num1 int, num2 int) (string, int) {
	sum := num1 + num2
	result := strconv.Itoa(num1) + " + " + strconv.Itoa(num2) + " = " + strconv.Itoa(sum)
	return result, sum
}

//可变参数,当做数组用
func add(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}

//返回值可以提前初始化
func add2(args ...int) (sum int) {
	sum = 0
	for _, value := range args {
		sum += value
	}
	return
}

//函数可以当形参用
func show(a int, b int, fun func(int, int) (string, int)) {
	fmt.Println(fun(a, b))
}

func main() {
	result, sum := addition(1, 1)
	fmt.Println(result)
	fmt.Println("结果:", sum)

	r, _ := addition(2, 2) //如果有返回值不接收,那可以用_代替
	fmt.Println(r)

	fmt.Println(add(1, 2, 3, 4, 5))

	fc := addition
	fmt.Printf("f的数据类型是: %T\n", fc)
	show(3, 3, fc) //函数可以当形参用

	//匿名函数,即没有名字的函数,一般定义完就使用,只使用一次
	func(data int) {
		fmt.Println("匿名函数调用", data)
	}(100) //这里直接进行调用
}

//初始化函数
func init() {
	fmt.Println("初始化函数,在执行main前会调用")
}
