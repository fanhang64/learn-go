package main

import "fmt"


func main()  {
	var flag = false

	for i:=0; i < 10; i++{
		for j := 'A'; j < 'Z'; j++{
			if j == 'C'{
				flag = true
				break
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag{
			break
		}
	}

	// goto语句, 少用
	for i:=0; i < 10; i++{
		for j := 'A'; j < 'Z'; j++{
			if j == 'C'{
				goto xx
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
	xx:  // label标签
		fmt.Println("end")


}