package main   // main包生成执行文件， 其他名字作为模块

import "fmt"   // 导入语句

// 函数外只能放置标识符（变量，常量，函数，类型）的声明
// fmt.Println("xxxxx")  // error 只能放到函数里

// 程序的入口函数
func main()  {
	fmt.Println("Hello world")
}