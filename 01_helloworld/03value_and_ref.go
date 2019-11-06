package main

import . "fmt"
import "reflect"


func main(){
	// 值和引用
	// 值类型：基本数据类型，int,float,bool,string,以及数组和struct
	// 特点：变量直接存储值(变量直接指向存在内存中的值)，内存通常在栈中分配，栈在函数调用完会被释放, 
	
	Println("hello")
	my_int := 123
	Println(&my_int)  // 0xc00000a0f0
	
	my_int = 456
	Println(&my_int)  //  0xc00000a0f0 指向相同地址 只是值拷贝  。值类型的变量的值存储在栈中

	var c = [3]int{1, 2, 3}  // 定义长度为3的int类型数组
	d := c
	Println(d, c)  // [1 2 3]  [1 2 3]
	
	d[1] = 100  // 将d中索引为1的值修改为100
	Println(d, c)  // [1 100 3]  [1 2 3]
	Println(&d, &c)  // 地址不同 
	Printf("%v, %p \n%v, %p\n", d, &d, c, &c)  //  [1 100 3], 0xc000064160  // [1 2 3], 0xc000064140
	Println(reflect.TypeOf(d))  // [3]int
	

	// 引用类型：指针，slice，map，chan等都是引用类型
	// 变量直接存放的就是一个内存地址值，这个地址值指向的空间存的才是值。所以修改其中一个，另外一个也会修改（同一个内存地址）
	// 引用类型必须申请内存才可以使用，make()是给引用类型申请内存空间。

	// TODO ...

}