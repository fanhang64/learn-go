package main

import "fmt"

func main(){
	b1 := true

	var b2 bool  // 默认false

	fmt.Println(b1)
	fmt.Printf("%T value: %v", b2, b2)  // %v 打印变量的值 无论任何类型

	b3 := 10  
	
	// b4 := bool(b3) // error 不能将整型转为布尔

	// b5 := b2 + b3  // error 不能进行数值运算
	
	fmt.Println(b3)
}
