package main

import "fmt"

func main()  {

	// 对于英文字符叫byte类型
	// 对于中文其他字符叫rune类型

	// 字符类型
	c1 := "哈"
	c2 := '哈'  // rune类型 也是(int32) 类型
	fmt.Printf("c1: %T, c2: %T\n", c1, c2)  // c1: string, c2: int32  

	c3 := "H"
	c4 := 'H'  // 默认是rune，也是int32
	c5 := byte('H')  // 强制转换为byte类型，也是uint8类型
	fmt.Printf("c3: %T, c4:%T, c5:%T\n", c3, c4, c5)  // c3: string, c4:int32


	// 字符串的遍历
	s := "hello世界"
	fmt.Println(s)

	// len() 获取的是byte字节的数量
	length := len(s)  // utf-8下一个汉字占3-4个字节

	fmt.Println(length)

	for i := 0; i < len(s); i++ {  // 使用byte类型遍历字符串
		// fmt.Println(s[i])  //   打印字节
		fmt.Printf("%v(%c) ", s[i], s[i])  // 打印字符, 可以打印出应为，中文出现乱码
	
	}

	for _, i := range s {  // 按照rune类型从字符串中拿出具体的字符 
		fmt.Printf("%c\n", i)  
	}

	// 字符串的修改
	s2 := "哈哈哒"  // 表示 '哈' '哈' '哒' 三个字符
	// s2[0] = "吧"  // error 字符串不能修改
	
	s3 := []rune(s2)  // 将s2强制转为rune切片

	// s3[0] = "把"  // error 
	s3[0] = '把'  // ok [哈 哈 哒] s[0]= ''应该是字符
	fmt.Println(string(s3))  // 将切片转为字符串


	// 类型转换
	// 注意： bool不能和其他类型转换

	// 整数转浮点数
	m := 10
	var f float64
	f = float64(m)
	fmt.Printf("%T, %f\n", f, f)  // float64, 10.000000

	// 浮点数转整数  
	// 小数部分被丢弃
	var v1 float32 = 0.999999
	fmt.Println(int(v1))  // 0
	v1 = -0.999999
	fmt.Println(int(v1))  // 0

}