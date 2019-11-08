package main

import (
	"flag"
	"fmt"
)

func main() {
	// 1)
	var aVar = 10
	a := aVar == 5 // false
	b := aVar == 10
	c := aVar != 5
	d := aVar != 10

	fmt.Println(a, b, c, d)  // false true true false


	// 2） 只有两个相同类型的值才可以进行比较
	//var q int = 10
	//var w float64 = 22.22

	//a = q == w  // errors 只有相同类型才可以比较  invalid operation: q == w (mismatched types int and float64)
	//fmt.Println(a)  //

	const cc = 123
	var dd = 33

	a = cc == dd  // ok 可以一个为常量 一个为变量，但类型必须相同
	fmt.Println(a)

	//a = 12 == true  // error 类型错误， 必须为相同int或者相同bool
	//fmt.Println(0 == false)  // error   类型不同


	// 布尔类型不会自动转为数字0,1 需要显式转换
	c = itob(10)
	fmt.Println(c)
	ddd := btoi(false)
	fmt.Println(ddd)

	// 布尔不可以进行与其他类型的强制类型转换
	//var n bool = false
	//fmt.Println(int(n)) // error  cannot convert n (type bool) to type int

	//fmt.Println(int(false)) // error cannot convert n (untyped bool) to type int

	fmt.Println(false + 123)  //  布尔不能进行数值计算
}

// 布尔转数字
func btoi(b bool)int {
	if b{
		return 1
	}
	return 0
}

// 数字转布尔
func itob(i int) bool {
	return i != 0
}
