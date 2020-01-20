package main

import "fmt"

func main()  {
	// 当i=5时跳出for循环
	for i:= 0; i < 10; i++{
		if i == 5{
			break  // 跳出循环
		}
		fmt.Println(i)
	}

	// break添加标签,
	// break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块
	haha:
	for x := 0; x < 11; x++{
		if x == 3{
			break haha
		}
		fmt.Println("haha label...", x)
	}

	
	// 当i=5时跳过此次for循环
	// continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。
	for i:=0; i < 10; i++{
		if i == 5{
			continue  //继续下一次
		}
		fmt.Println(i)
	}

	wawa:
	for x := 0; x < 11; x++{
		if x == 3{
			continue wawa
		}
		fmt.Println("haha label...", x)
	}

}