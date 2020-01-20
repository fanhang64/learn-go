package main

import "fmt"


func main()  {
	age := 19

	// 基本格式
	if age > 18 {
		fmt.Println("大于18碎了")
	}else{
		fmt.Println("小于18岁")
	}

	// Go语言规定与if匹配的左括号{必须与if和表达式放在同一行，{放在其他位置会触发编译错误。 同理，与else匹配的{也必须与else写在同一行，else也必须与上一个if或else if右边的大括号在同一行。
	// if age > 1  // error 
	// {

	// }

	age = 65
	// 多个判断条件
	if age > 35{
		fmt.Println("年龄大于35")
	}else if age > 18 {
		fmt.Println("大于18, 小于35")
	}else{
		fmt.Println("小于18")
	}

	// if 特殊写法
	// ageNow作用域此时只在if条件判断语句中
	if ageNow:=20; ageNow > 18{
		fmt.Println("年龄大于18")
	}else{
		fmt.Println("年龄小于18")
	}
	// fmt.Println(ageNow)  // error undefined
	
}