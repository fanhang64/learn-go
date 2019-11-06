package main

import "fmt"

func main(){
	//  1) 变量
	// 变量名由数字字母下划线组成,第一个不能是数字
	//  var 变量名 类型
	//  var 变量名1, 变量名2 类型

	var a string = "my string1111"
	fmt.Println(a)

	var b, c, d, e int = 1, 2, 3, 4
	fmt.Println(b, c, d, e)  // 1 2 3 4 
	
	// 变量声明 但不初始化
	// 数值类型（包括complex）返回0 
	var v_name int
	fmt.Println(v_name)  // 0
	var v_name3 float32
	fmt.Println(v_name3)  // 0

	var v_name5 [2]int
	fmt.Println(v_name5)  // [0 0]

	// 布尔类型 返回false
	var v_name4 bool
	fmt.Println(v_name4)  // false

	// 字符串返回""
	// var v_name string  // error v_name 重定义
	var v_name2 string
	fmt.Println(v_name2 + "=====")  // ""

	// 以下返回nil
		//  var a * int
		//  var a []int
		//  var a map[string] int 
		//  var a chan int
		//  var a func(string) int 
		//  var a error // error接口
	
	//  2)  根据值自行判断类型
	var var_name = false
	fmt.Println(var_name)
	var var_name1 = 123
	fmt.Println(var_name1)

	// 3) 省略var  := 左侧如果没有声明新的变量（只要有一个是新的就不会报错），就产生编译错误
	var ccc int
	// ccc := 123  // error  no new variables on left side of :=
	ccc, ddd := 123, 456  
	// fmt.Println(ccc)  // error  := 是一个声明语句 声明的ddd必须使用 不使用会报错
	fmt.Println(ccc, ddd)  // ok

	var eee, ggg int
	eee, ggg, hh := 1, 2, 3
	fmt.Println(eee, ggg, hh)

	fff := "hello world"  // 等同于 var s struing = "hello world"
	fmt.Println(fff)

	fmt.Println(a11, "=-=====")   // 0 =-=====
}

// 一般用于声明全局变量
var (
	a11 int 
	b float32
)

// asd, cas := 111, 222   // 只能在函数内声明
