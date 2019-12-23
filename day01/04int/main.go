package main

import "fmt"


// 整型

func main()  {
	var a = 10  // 类型推导

	fmt.Printf("%d\n", a)

	// 八进制
	b := 077  
	fmt.Printf("%d %o\n", b, b)

	// 十六进制数
	c := 0xff
	fmt.Printf("%d, %x\n", c, c)
	
	// 查看变量的类型
	fmt.Printf("%T\n", c)  // 默认int类型

	// 声明一个int8类型的变量
	d := int8(9)
	fmt.Printf("%T\n", d)  // 默认int8类型

}