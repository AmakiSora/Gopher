package main

import "fmt"

/*
	程序流程控制
	分支控制 if-else switch
	循环控制 while do-while for
*/
func main() {
	/*
		if-else
		在Go语言中,if后面的{}是必须有的,就算你只写一行
	*/
	if i := 1; i == 0 { //Go语言支持在if中定义变量
		fmt.Println("i == 0")
	} else if i > 1 { //括号()可写可不写
		fmt.Println("i > 0")
	} else {
		fmt.Println("other")
	}

	//if true {} //死循环写法

	/*
		switch
		在Go语言中,匹配项后面不需要加break,当执行完后自动退出switch
	*/
	k := 233
	switch key := 1; key {
	case 'a', 'b', 3: //可以有多个匹配值
		fmt.Println("ab3")
	case '1':
		fmt.Println("111")
	case 1, k: //可以是任意表达式
		fmt.Println("1k")
	default:
		fmt.Println("other")
	}

	switch { //switch后可以不带表达式
	case k == 23:
		fmt.Println("k == 23")
	case k == 233:
		fmt.Println("k == 233")
		fallthrough //穿透一层,即输出k == 233后,继续执行下一层case
	case k == 2:
		fmt.Println("k == 2")
	case k == 3:
		fmt.Println("k == 2")
	}

	var x interface{}
	var y = 10.0
	x = y
	switch x.(type) { //type-switch,用于判断某个interface变量中实际指向的变量类型
	case nil:
		fmt.Println("")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case bool, string:
		fmt.Println("bool or string")
	}

	/*
		for
		go语言中只有for这一种循环
	*/
	for i := 0; i < 5; i++ { //可以不写括号()
		fmt.Println("五次循环")
	}

	i := 0
	for i < 3 {
		fmt.Println("三次循环")
		i++
	}

	for {
		fmt.Println("死循环")
		if i < 6 {
			fmt.Println("跳过本次循环")
			i += 2
			continue //跳过本次循环,开始下次循环
		}
		if i >= 10 {
			fmt.Println("结束循环")
			break //跳出循环
		}
		i++
	}

	//for range
	str := "hello,一条字符串"
	for i, value := range str { //i为索引,value为具体值
		fmt.Printf("%d      %c\n", i, value) //注意:中文占三个字节,所以i索引会相隔3个
	}

	//goto 无条件跳转到指定位置执行(不建议使用)
	fmt.Printf("1 ")
	fmt.Printf("2 ")
	fmt.Printf("3 ")
	goto l1
	fmt.Printf("4 ")
	fmt.Printf("5 ")
	fmt.Printf("6 ")
l1:
	fmt.Printf("7 ")
	fmt.Printf("8 ")
	fmt.Printf("9 ")

}
