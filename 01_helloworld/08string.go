package main

import (
	"fmt"
)


func main(){
	// 一个字符串是一个不可改变的byte序列， 字符串占用大小（当为ascii表的字符时占用1个字节，当为其他字符时占用2-4个字节）
    // 字符串是字节的数组
   	// 使用"" 定义字符串

   	var str string = "我是go语言"

   	fmt.Println(str)

   	// len(str) 用来获取所占字节长度或者ascii字数
   	fmt.Println(len(str))  // 14
   	fmt.Println(len("a"))  // 占用1个字节
   	fmt.Println(len("go"))  // 占用2个字节
   	fmt.Println(len("我"))  // 占用3个字节

   	fmt.Println(str[0])  // 第一个字节
   	fmt.Println(str[6])  // 103   ascii中g的位置

   	//fmt.Println(&str[1])  // error不能获取某个字节的地址
	fmt.Println(&str)  // ok



	// 2) 字符串拼接 +
	myStr := "Beginning of the string " +
		"second part of string" +
		" third part of string"
	fmt.Println(myStr)

	myStr += " fourth part of string"  // 还可以使用+=
	fmt.Println(myStr)

	// 3) 定义多行
	// `` 反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出
	myString := `  
第一行
第二行
第三行
\r\n
\t
第四行`
	fmt.Println(myString)
}
