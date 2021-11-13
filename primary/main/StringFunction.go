package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	string常用函数
*/
func main() {
	str := "go语言"

	//统计字符串长度(按字节计算,utf-P中汉字占3字节)
	fmt.Println(len(str))

	//字符串遍历(方式1)
	for i, value := range str {
		fmt.Printf("索引:%d,值:%c\n", i, value)
	}

	//字符串遍历(方式2)
	c := []rune(str)
	for i := 0; i < len(c); i++ {
		fmt.Printf("值:%c\n", c[i])
	}

	//字符串转整数
	num, _ := strconv.Atoi("666")
	fmt.Println(num)

	//整数转字符串

	fmt.Println(strconv.Itoa(233))

	//统计一个字符串中有几个子串
	fmt.Println(strings.Count("上上下下左右左右baba", "ba"))

	//不区分大小写的字符串比较
	fmt.Println(strings.EqualFold("go", "Go"))

	//区分大小写的字符串比较
	fmt.Println("go" == "Go")

	//返回子串在字符串中第一次出现的索引值,如果没有则返回-1
	fmt.Println(strings.Index("996996", "6"))

	//字符串的替换(原字符串,被替换的子串,替换的子串,替换个数-1为全换)
	fmt.Println(strings.Replace("233233233233", "233", "666", 3))

	//按照指定的字符进行切割,将字符串拆分成字符串数组
	fmt.Println(strings.Split("哈-啊-嗯", "-"))

	//字符串字母大小写转化
	cosmos := "Amaki SorA"
	fmt.Println(strings.ToLower(cosmos)) //转小写
	fmt.Println(strings.ToUpper(cosmos)) //转大写
	fmt.Println(strings.Title(cosmos))   //首字母大写

	//去除字符串左右两边空格
	fmt.Println(strings.TrimSpace("         ? ? ?         "))

	//去除字符串左右两边指定字符
	fmt.Println(strings.TrimLeft("~~哦~吼~~", "~"))  //去除左边
	fmt.Println(strings.TrimRight("~~哦~吼~~", "~")) //去除右边
	fmt.Println(strings.Trim("~~哦~吼~~", "~"))      //去除两边

	//判断字符串是否以指定字符串开头
	fmt.Println(strings.HasPrefix("https://996.com", "https"))

	//判断字符串是否以指定字符串结尾
	fmt.Println(strings.HasSuffix("https://996.com", "com"))

}
