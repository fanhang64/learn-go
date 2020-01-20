package main

import "fmt"


func main()  {
	var n = 5

	// 简化大量判断
	switch n {
	case 1: fmt.Println("11")
	case 2: fmt.Println("22")
	case 3: fmt.Println("33")
	case 4: fmt.Println("44")
	case 5: fmt.Println("55")
	default:
		fmt.Println("默认")
		
	}

	// 变种1
	switch n:=3;n {
	case 1:
		fmt.Println("111")		
	case 3:
		fmt.Println("333")
	default:
		fmt.Println("moren")
	}

	// 变种2, case多个条件
	switch n:=3;n {
	case 1,3,5,7,9:
		fmt.Println("奇数")
	case 2,4,6,8,0:
		fmt.Println("偶数")
		
	}

	// 变种3，case判断条件
	age := 19
	switch {  // 此时switch没有表达式，如果添加表达式case 条件会报error
	case age > 18:
			fmt.Println("大于18")
	case age < 18:
			fmt.Println("小于18")
	case age == 18:
			fmt.Println("等于18")
	}

	// 变种4： 兼容c语言穿透，C语言默认需要加break防止穿透。go相反不需要加，如果想同C语言一样，可以使用fallthrough
	s := 'a'
	switch  {
	case s == 'a':
		fmt.Println("a")
		fallthrough       // 穿透下一个case
	case s == 'b':
		fmt.Println("b")
	case s == 'c':
		fmt.Println("c")
	}
	
}