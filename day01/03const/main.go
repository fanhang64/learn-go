package main

import "fmt"


// const 常量
// 定义了常量之后不能修改，在程序运行期间不会改变
const pi = 3.1415926

// 批量声明常量
const (
	STATUSOK = 200
	NOTFOUND = 404
)

// 批量声明常量值，如果某一行声明后，没有赋值，同上一行
// n2 n3 默认和n1相同
const (
	n1 = 100
	n2
	n3

)

const (
	m1 = 100
	m2
	m3
	m4 = 200
	m5

)


// iota 是go语言中常量计数器, 只能在常量的表达式中使用
// 在const中 被重置为0
// const中每新增一行常量声明将使iota计数一次
const (
	a1 = iota  // 0 
	a2         // 1  相当于 a2 = iota
	a3         // 2  相当于 a3 = iota
	a4 = iota  // 3  
)


// iota常见示例：
// 1） 使用_跳过某些值
const (
	b1 = iota  // 0
	b2         // 1
	_          // 忽略掉2
	b3         // 3
)

// 2）插队
const (
	c1 = iota   // 0
	c2 = 100    // 100   const中每新增一行常量声明将使iota计数一次  相当于1的位置
	c3 = iota   // 2
	c4          // 3   == c4 = iota
)

// 3） 多个iota声明在一行
const (
	d1, d2 = iota+1, iota +2   // 0+1 0+2
	d3, d4 = iota +1, iota +2  // 1+1 1+2
	d5, d6                     // 2+1 2+2 同上 d5,d6 = = iota +1, iota +2 
)

// 4） 定义数量级

const (
	_ = iota
	KB = 1 << (10 * iota)  // 1左移10位即1024
	MB = 1 << (10 * iota)  // 1左移10 *2 位
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
func main()  {

	// pi = 123  // 修改常量 error
	fmt.Println(n1, n2, n3)  // 100 100 100
	fmt.Println(m1, m2, m3, m4, m5)  // 100 100 100  200 200
	iota := 1000   // ok 
	
	fmt.Println(iota)  // 1000

	fmt.Println(a1, a2, a3, a4)  // 0 1 2 3
	fmt.Println(b1, b2, b3)  // 0 1 3

	fmt.Println(c1, c2, c3, c4)  // 0 100 2 3
	fmt.Println(d1, d2, d3, d4)  // 1 2 2 3

	fmt.Println(KB, MB, GB, TB, PB)
	
}