package main

import "fmt"

/*
	切片
	Go 语言切片是对数组的抽象,是一种特有的数据类型
	Go 数组的长度不可改变,Go中提供了一种内置类型切片("动态数组"),与数组相比切片的长度是不固定的,可以追加元素,在追加时可能使切片的容量增大
*/
func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}

	//定义一个切片,因为切片是动态变化的,所以[]里不写
	var sl []int

	//切片构建在数组之上
	sl = arr[1:4]
	fmt.Println(sl)

	//切片元素个数
	fmt.Println(len(sl))

	//切片的容量(会动态变化)
	fmt.Println(cap(sl))

	//切片的内存地址为切的数组的内存地址
	fmt.Println(&sl[0])
	fmt.Println(&arr[1])

	//改变切片的值也会改变被切的数组的值
	sl[0] = 233
	fmt.Println(arr[1])

	//make方法定义切片,make底层会创建一个数组,对外不可见,不能直接操作这个数组
	slm := make([]int, 2, 5)
	fmt.Println(slm)
	fmt.Println(len(slm))
	fmt.Println(cap(slm))

	//切片可以扩容,扩容后,如果长度大于容量,则底层地址会变为新数组的地址(类似java的arraylist)
	fmt.Println(&slm[0])
	sla := append(slm, 77, 88, 99, 00)
	fmt.Println(sla)
	fmt.Println(&sla[0])

	//切片也可以追加切片
	sla = append(sla, sla...)
	fmt.Println(sla)

	//切片的拷贝
	slx := []int{1, 2, 3, 4, 5}
	sly := make([]int, 2)
	copy(slx, sly)
	fmt.Println(slx)

}
