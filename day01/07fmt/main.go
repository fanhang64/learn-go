package main

import "fmt"

func main()  {
	var n = 100

	// fmt包总结
	fmt.Printf("%T\n", n)  // 查看类型
	fmt.Printf("%v\n", n)  // 打印变量的值
	fmt.Printf("%b\n", n)   
	fmt.Printf("%o\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%x\n", n)

	// 字符串
	var s = "hello"

	fmt.Printf("%s\n", s)  // hello
	fmt.Printf("%v\n", s)  // hello
	fmt.Printf("%#v\n", s)  // "hello"  打印字符串
}