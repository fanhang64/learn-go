package main

import "fmt"

// 在函数体外声明的变量称之为全局变量
// 全局变量与局部变量名称可以相同，但是函数体内的局部变量会被优先考虑。
var c int  = 100


func main()  {
	// 局部变量的作用域只在函数体内，参数和返回值变量也是局部变量
	var a int = 3 
	var b int = 4

	// c := a + b 
	fmt.Printf("a = %d, b= %d, c= %d\n", a, b, c)

	// 形参
	d := sum(a, b)
	fmt.Println(d,)
}


// 形参也是局部变量
func sum(a, b int) int {
	num := a + b
	return num
}
