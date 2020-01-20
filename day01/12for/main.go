package main

import "fmt"


func main() {
	// 基本格式
	// 
	for i:= 0; i < 10; i++{
		fmt.Println(i)
	}
	// fmt.Println(i)  // error 找不到i

	// 变种1： 初始语句可以省略(在for循环外定义)
	var j = 4
	for ; j < 10; j++{
		fmt.Println(j)
	}
	fmt.Println(j)  // 10
	// 变种2： 结束语句也可省略
	var x = 5
	for ; x < 10; {
		fmt.Println(x)
		x++
	}
	// 同变种2
	x = 5
	for x < 10{
		fmt.Println(x)
		x++
	}

	// 无限循环
	// for {
	// 	fmt.Println("1")
	// }
	
	// for range循环

	s := "hello世界"
	for i,v := range s{  // i索引 v字符
		fmt.Println(i, v)
		fmt.Printf("%v(%c)\n", v, v)  
	}
}