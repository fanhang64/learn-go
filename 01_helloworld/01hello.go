package main

import "fmt"

func main(){
	// { 注意 {不能单独一行
	fmt.Println("hello world")

	// 基本语法
	// （1） 行分隔符  在go语言中 一行带边一个语句结束，每个语句结束不需要加; 一行多个语句需要加;
	//          
	fmt.Println("hang fen ge "); fmt.Println("  符号");  // ok
	// （2） 注释
	//  // 单行注释
		/* 多行注释 */
	//  （3） 标识符
	// 一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字
	
	// （4） 字符串连接
	fmt.Println("hello" + "world! go")

	// （5） 数据类型
	// 主要包括四大类： 布尔类型， 数字类型， 字符串类型，派生类型
	//   布尔类型： 只能是false和true两种

	var b bool = true
	fmt.Println(b)
	b = false
	fmt.Println(b)

	//  数字类型： 整型 浮点型float32 和浮点型64  支持复数
		// 无符号整型  uint8 uint16 uint32 uint64 
		// 有符号整型  int8 int16 int32 int64
		var c int64 = 11111
		fmt.Println(c)

		// 浮点数
		// float32 float64 
		// complex64 32 位实数和虚数
		//  complex128  64 位实数和虚数
		var myFloat float32 = 1111.111
		fmt.Println(myFloat)
		var myComplex complex128 = 11111.111
		fmt.Println(myComplex)

	//  字符串类型： 字符串是由单个字节连接起来的，默认使用UTF-8编码
	//  派生类型：
		// 1. 指针类型Pointer
		// 2.  数组类型
		// 3. 结构化类型
		// 4. channel类型
		// 5. 函数类型
		// 6. 切片类型
		// 7. 接口类型(interface)
		// 8. Map类型


}

