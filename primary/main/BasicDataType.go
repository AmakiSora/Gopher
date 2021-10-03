package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

/*
	基本数据类型
		整数型:int,int8,int16,int32,int64,uint,byte
		浮点型:float32,float64
		布尔型(bool):布尔型只有true和false,其他数字等都不行
		字符串(string):在go里字符串算基本数据类型
	注意:go语言没有字符型,使用byte来保存单个字符
*/
func main() {
	/*
		整数型		符号		占用空间	范围(64位系统)
		int(默认)	有符号	8字节	-9223372036854775808 ~ 9223372036854775807
		int8		有符号	1字节	-128 ~ 127
		int16		有符号	2字节	-32768 ~ 32767
		int32		有符号	4字节	-2147483648 ~ 2147483647
		int64		有符号	8字节	-9223372036854775808 ~ 9223372036854775807
		uint		无符号	8字节	0 ~ 18446744073709551615
		uint8		无符号	1字节	0 ~ 255
		uint16		无符号	2字节	0 ~ 65535
		uint32		无符号	4字节	0 ~ 4294967295
		uint64		无符号	8字节	0 ~ 18446744073709551615
		rune		有符号	4字节	-2147483648 ~ 2147483647	等价int32,表示一个Unicode码
		byte		无符号	1字节	0 ~ 255						本质上是uint8,存储字符使用
	*/
	var i int8 //默认为0
	fmt.Println("i的值:", i)
	fmt.Printf("i的类型:%T\n", i)                  //查看数据类型
	fmt.Printf("i的占用空间:%d\n", unsafe.Sizeof(i)) //查看占用空间
	fmt.Println("-----------------------------------")
	/*
		浮点型		占用空间		范围(64位系统)
		float32		4字节		-3.403E38 ~ 3.403E38
		float64(默认)8字节		-1.798E308 ~ 1.798E308
	*/
	var f = .123 //默认为0
	fmt.Println("f的值:", f)
	fmt.Printf("f的类型:%T\n", f)
	fmt.Printf("f的占用空间:%d\n", unsafe.Sizeof(f)) //查看占用空间
	fmt.Println("1.11e3 ->", 1.11e3)            //科学计数法,等价于1.11 * 10的3次方
	fmt.Println("-----------------------------------")
	/*
		布尔类型(占1字节)
		Go语言里只允许取值true和false,其他数值(0和非0)不允许
	*/
	var b = true //默认为false
	fmt.Println("b的值:", b)
	fmt.Printf("b的类型:%T\n", b)
	fmt.Println("-----------------------------------")
	/*
		字符串
		Go的单个字符一般由byte来存储
		Go的字符串是有字节组成的(不是字符)
		Go语言的字符使用UTF-8编码
		双引号("")会识别转义字符
		反引号(``)以字符串原始形式输出,不会识别任何字符
	*/
	var c byte = 'a' //默认为""
	fmt.Println("c的值:", c)
	fmt.Printf("c的类型:%T\n", c)
	str := "字符串"
	fmt.Println("str的值:", str)
	fmt.Printf("str的类型:%T\n", str)
	fmt.Printf(`str的类型:%T\n`, str) //反引号不会将\n翻译成换行
	fmt.Println()
	fmt.Println("-----------------------------------")
	/*
		数据类型转换
		和其他语言不一样,Go在不同类型的变量之间赋值需要显式转换(不能自动转换)
		范围小的可以转成范围大的,范围大的也可以转成范围小的,只是会溢出,得到的结果不符合预期
	*/
	var t1 int = 9999
	var t2 float32 = float32(t1)
	var t3 float64 = float64(t2) //同样是float也要显式转换
	var t4 int8 = int8(t3)       //范围大变范围小,会溢出
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(t4)
	fmt.Println("-----------------------------------")
	/*
		其他数据类型转成string
		方式1 fmt.Sprintf("%参数",表达式)
		方式2 strconv函数
	*/
	var num1 = 233.233 //方式1
	var toStr = fmt.Sprintf("%f", num1)
	fmt.Printf("%T,%v\n", toStr, toStr) //类型,值

	var num3 = 233                      //方式2
	toStr = strconv.Itoa(num3)          //int -> string
	fmt.Printf("%T,%v\n", toStr, toStr) //类型,值

	var num2 = 233.456                             //方式2
	toStr = strconv.FormatFloat(num2, 'f', 10, 64) //float64 -> string
	//strconv.FormatFloat(变量,格式,保留的位数,类型)
	fmt.Printf("%T,%v\n", toStr, toStr) //类型,值
	fmt.Println("-----------------------------------")

	/*
		string转其他数据类型
		strconv函数
		如果错误,返回类型的默认值
	*/
	var str1 = "true"
	var B bool
	B, _ = strconv.ParseBool(str1) //_表示忽略,后面函数章节会解释
	fmt.Printf("%T,%v\n", B, B)    //类型,值

	var str2 = "996"
	var n int64
	n, _ = strconv.ParseInt(str2, 16, 64)
	//strconv.ParseInt(变量,进制,类型)
	fmt.Printf("%T,%v\n", n, n) //类型,值
	fmt.Println("-----------------------------------")

	/*
		转义字符
		%v	值的默认格式表示
		%+v	类似%v，但输出结构体时会添加字段名
		%#v	值的Go语法表示
		%T	值的类型的Go语法表示
		%%	百分号
		%t	单词true或false
		%b	表示为二进制
		%c	该值对应的unicode码值
		%d	表示为十进制
		%o	表示为八进制
		%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
		%x	表示为十六进制，使用a-f
		%X	表示为十六进制，使用A-F
		%U	表示为Unicode格式：U+1234，等价于"U+%04X"
		%b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
		%e	科学计数法，如-1234.456e+78
		%E	科学计数法，如-1234.456E+78
		%f	有小数部分但无指数部分，如123.456
		%f:    默认宽度，默认精度
		%9f    宽度9，默认精度
		%.2f   默认宽度，精度2
		%9.2f  宽度9，精度2
		%9.f   宽度9，精度0
		%F	等价于%f
		%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
		%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
		%s	直接输出字符串或者[]byte
		%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
		%x	每个字节用两字符十六进制数表示（使用a-f）
		%X	每个字节用两字符十六进制数表示（使用A-F）
		%p	表示为十六进制，并加上前导的0x
		'+'	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
		' '	对数值，正数前加空格而负数前加负号；
		'-'	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
		'#'	切换格式：
		  	八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）；
		 	对%q（%#q），如果strconv.CanBackquote返回真会输出反引号括起来的未转义字符串；
		 	对%U（%#U），输出Unicode格式后，如字符可打印，还会输出空格和单引号括起来的go字面值；
		  	对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格；
		'0'	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；



	*/

}
