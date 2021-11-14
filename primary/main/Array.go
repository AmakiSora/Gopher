package main

import "fmt"

/*
	数组
	注意:go语言的数组传递是值传递,不是引用传递
*/
func main() {
	//数组的定义,不赋值默认为0
	var arr [5]int
	var twoArr [2][3]int //二维数组

	//初始化
	arr = [5]int{3, 4, 5, 6, 7}
	var arr2 = [...]int{2, 3, 3}
	arr3 := [...]int{2: 233, 6: 666} //{位置:值}
	twoArr = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	//赋值
	arr[0] = arr2[0]

	//遍历
	for i, value := range arr3 {
		fmt.Println(arr3[i])
		fmt.Println(value)
	}
	fmt.Println(twoArr)

	//数组的数据类型(长度也是类型的一部分)
	fmt.Printf("数组的数据类型:%T\n", arr)

	//数组传递为值传递
	changeArr2(arr2)
	fmt.Println(arr2)

	//数组的引用传递可以利用指针实现
	changeArr2Pointer(&arr2)
	fmt.Println(arr2)
}

//数组的值传递
func changeArr2(arr [3]int) {
	arr[0] = 99999
}

//数组的引用传递
func changeArr2Pointer(arr *[3]int) {
	(*arr)[0] = 66666
}
