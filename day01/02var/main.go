package main

import "fmt"

// GO中使用驼峰命名法

// 声明变量
// var name string
// var age int
// var isOk bool

// 批量声明
var (
	name string   // 默认值 ""
	age int       // 默认值 0
	isOk bool     // 默认值 false
	gender bool
)

func main()  {
	name = "fanhang64"
	age = 16 // 变量赋值
	isOk = true 

	// go语言中函数内的变量声明必须使用，不使用编译不过去
	// 全局变量可以声明不使用
	// var haha string  // 声明 不使用error
	//  gender 全局变量声明  不使用ok

	fmt.Println(age)  // 打印指定变量，结尾添加换行
	fmt.Printf("name:%s\n", name)  // 占位符
	fmt.Print(isOk)  // 输入isok内容在终端中
	fmt.Println()  // 空行

	var s1 string = "fhz"  // 声明同时赋值
	fmt.Println(s1)

	// 类型推导（根据值判断类型）
	var s2 = "20"
	fmt.Println(s2)  // ok 

	// 短变量声明
	// 只能在函数里用
	s3 := "hahah"
	fmt.Println(s3)

	// s3 := "error "  // 同一个作用域{}中不能重复声明变量

	// 匿名变量  _ 可以忽略返回值，不占用内存，可以重复声明
}