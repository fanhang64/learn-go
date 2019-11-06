package main

import "fmt"


func GetData() (int, int) {
	return 1, 2
}

func main()  {
	// 匿名变量 使用 "_" ,不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算

	a, _ := GetData()
	// fmt.Println(a, _)  // error 不能使用_
	_ = 111  // ok 可以赋值
	// a = _  // 不能赋值

	fmt.Println(a)

	// 注意： 匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。


}