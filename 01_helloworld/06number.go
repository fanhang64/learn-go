package main

import "fmt"


func main()  {
	// 数值类型
	
	// (1) 整型
	//  按长度分为：int8、int16、int32、int64 
	//  对应的无符号整型：uint8、uint16、uint32、uint64

	var a1 = 10  // 同 var a int 等于cpu的位数
	var a int = 10
	fmt.Println(a1)
	fmt.Printf("%d\n", a)  // 10
	fmt.Printf("%b\n", a)  // 1010 二进制
	fmt.Printf("%o\n", a)  // 12 八进制输出
	fmt.Printf("%d,,,,, %p\n", a, &a)

	var b int = 077  // 八进制以0开头
	fmt.Printf("%o\n", b) // 77
	fmt.Printf("%d\n", b) // 十进制输出   66 
	
	var c int = 0xff  // 十六进制0x
	fmt.Printf("%x\n", c)  // ff
	fmt.Printf("%X \n", c)  // FF

	var d int8 = 1
	var d2 int16 = 11
	var d3 int32 = 111
	var d4 int64 = 1111
	fmt.Printf("%d,,,,, %p\n", d, &d)
	fmt.Printf("%d,,,,, %p\n", d2, &d2)
	fmt.Printf("%d,,,,, %p\n", d3, &d3)
	fmt.Printf("%d,,,,, %p\n", d4, &d4)

	var d5 uint = 2   // 同于cpu的位数
	var d6 uint8 = 22
	var d7 uint16 = 222
	var d8 uint32 = 2222
	var d9 uint64 = 22222

	fmt.Printf("%d, %p,\n %d, %p,\n %d, %p,\n %d, %p,\n %d, %p\n", d5, &d5, d6, &d6, d7, &d7, d8, &d8, d9, &d9)


	// (2) 浮点数
	// 一个 float32 类型的浮点数可以提供大约 6 个十进制数的精度，而 float64 则可以提供约 15 个十进制数的精度，通常应该优先使用 float64 类型，因为 float32 类型的累计计算误差很容易扩散，并且 float32 能精确表示的正整数最大值有限

	var f0 = 1.11  // 同 var fo float64 = 1.11

	var f1 float32 = 1
	var f2 float64 = 11

	fmt.Printf("%f, %p, %f, %p\n", f1, &f1, f2, &f2)
	fmt.Println(f0)

	// 声明时可以只写整数部分或者小数部分
	var f3 float64 = 1.
	var f4 float64 = .123

	fmt.Printf("%f, %f\n", f3, f4)  // 1.000000, 0.123000
	
	// 科学计数法
	const f5 = 16.32e10

	// fmt.Printf("%d， %d", f5, 1.23E1235)  // constant 1.23e+1235 overflows float64 ????


	fmt.Printf("%.2f\n", f4) // 控制几位小数

	
	// (3) 复数
	var myComp3 = 1 + 2i // 等同于 var myComp3 complex128 = 1 + 2i
	fmt.Println(myComp3)

	var myComp complex64 = 123 + 1i
	fmt.Println(myComp) // (123+1i)
	myComp += 1
	fmt.Println(myComp) // (124+10i)

	myComp += 1 - 2i
	fmt.Println(myComp) // (125-1i)

	var myComp2 = complex64(1)
	fmt.Println(myComp2)  // (1+0i)


	fmt.Println(real(myComp2)) // 实部 1
	fmt.Println(imag(myComp2))  // 虚部 0
}