package main

import "fmt"
// import "math"

func main()  {
	// math.MaxFloat32  // 32位浮点数最大值
	// math.MaxFloat64  // 64位浮点数最大值

	f1 := 1.2345
	fmt.Printf("%T\n", f1)  // 默认go语言中浮点数是float64

	f2 := float32(1.23455)
	fmt.Printf("%T\n", f2)  // 显示声明float32
	
	// f1 = f2  // error 不能直接把float64传给float32

	// 复数
	var c1 complex64
	c1 = 1 + 2i
	
	var c2 complex128
	c2 = 2 + 3i

	fmt.Println(c1)
	fmt.Println(c2)
}