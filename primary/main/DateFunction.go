package main

import (
	"fmt"
	"time"
)

/*
	时间相关常用函数
*/

func main() {
	//获取当前时间
	now := time.Now()
	fmt.Printf("数据类型:%T\n现在时间为: %v\n", now, now)

	//使用结构体中的方法获取具体时间
	fmt.Printf("年: %v\n", now.Year())
	fmt.Printf("月: %v\n", now.Month())
	fmt.Printf("日: %v\n", now.Day())
	fmt.Printf("时: %v\n", now.Hour())
	fmt.Printf("分: %v\n", now.Minute())
	fmt.Printf("秒: %v\n", now.Second())
	fmt.Printf("星期: %v\n", now.Weekday())
	fmt.Printf("1月1号至今的天数: %v\n", now.YearDay())
	fmt.Printf("纳秒: %v\n", now.Nanosecond())
	fmt.Printf("协调世界时: %v\n", now.UTC())

	//日期格式化
	fmt.Printf(now.Format("2006-01-02 15:04:05")) //必须要填这个(就nm脑瘫)

	//
}
