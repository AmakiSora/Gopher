package main

import (
	"errors"
	"fmt"
)

/*
	错误
*/
func main() {

	catchError()

	err := customError(0)
	if err != nil {
		fmt.Println("错误为:", err)
		//使用panic中断程序,后面的代码不会执行
		panic(err)
	}
	fmt.Println("执行了")
}

//捕获错误
func catchError() {
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

//自定义错误
func customError(b int) error {
	a := 1
	if b == 0 {
		//抛出自定义错误
		return errors.New("除数不能为0!")
	} else {
		fmt.Println(a / b)
		//没有错误则返回nil
		return nil
	}
}
