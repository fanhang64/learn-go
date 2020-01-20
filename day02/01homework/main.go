package main

import "fmt"  // 导入的包也必须使用

// 单行注释
 
/* 多行注释 */

// go语言函数外的语句必须以关键字开头
var name = "撒按时"
var age int
// age = 18  // error age是变量名，不为关键字

// main函数是入口返回，必须没有参数也没有返回值
func main() {
	// 函数内部定义的变量必须使用
	var isOk = true 
	fmt.Println(isOk)

	// 变量
	// 变量声明方式
	// var name string
	// var name = "吃阿斯顿" // 根据初始值推断类型
	// 函数内部专属  name := "asdasd"
	// 匿名变量(哑元变量)
	// 当有些数据必须用数据接受，但又不使用它时，可以用_来接受
	
	// 常量
	const PI = 3.1415213
	const UserNotExistErr = 10000
	
	// iota实现枚举
	// 二个要点 
	// 在const关键字出现的时候被重置为0
	// 每新增一行 常量声明 iota累加1

	// 流程控制
	// if for

	// for range
	for i := range "helloworld"{
		fmt.Println(i)  // 从0开始索引
	}

	for i, v := range "hello 世界"{
		fmt.Printf("%d %c\n", i, v)
	}
	
	// 9*9乘法表
	for i:=1; i < 10; i++{
		for j := 1; j < i+1; j++{
			fmt.Printf("%d * %d = %d  ", i, j, i*j)
		}
		fmt.Println()
	}

	var cc =123
	fmt.Printf("%T\n", cc)

	// go语言没有办法直接定义2进制
	var xx = 0222   // 八进制
	var xx2 = 0xff  // 十六进制
	fmt.Println(xx, xx2)  // 直接打印为10进制
}

// 数据类型
// 整型: 默认int类型
//   无符号整型 uint8 uint16 uint32 uint64
//   带符号整型 int8 int16 int32 int64
//   int具体是32位还是64位根据操作系统判断
// uintptr 表示一指针

// 浮点型
//  float64  float32 默认float64

// 类型不同不能进行比较
// bool类型不能和其他类型转换

// 字符串总一个汉字一般占用3个字节


