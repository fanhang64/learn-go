package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// 1）字符串的统计 len() 和 RuneCountInString()
	c1 := "hello world"

	c2 := "你好世界"  //  Go 语言的字符串都以 UTF-8 格式保存，每个中文占用 3 个字节，因此使用 len() 获得两个中文文字对应的 6 个字节

	fmt.Println(len(c1), len(c2))  // 11   12

	// 统计字符utf-8个数
	fmt.Println(utf8.RuneCountInString(c1), utf8.RuneCountInString(c2))  // 11 4


	// 2) 对字符串的遍历
	// ASCII 字符串遍历直接使用下标。
	// Unicode 字符串遍历用 for range
	myString := "遍历字符串样例"

	// 便利每个ascii字符
	for x := 0; x < len(myString); x++{
		fmt.Printf("%c, %d\n", myString[x], myString[x])  // 乱码
	}

	// 正确的做法
	// 按照utf-8遍历
	for _, s := range myString{
		fmt.Printf("%c, %d\n", s, s)
	}


	// 3) 字符串的截取
	// 获取某个字符串的子串
	//strings.Index：正向搜索子字符串。
	//strings.LastIndex：反向搜索子字符串。
	//搜索的起始位置可以通过切片偏移制作。

	example := "呵呵哈哈，电死你来的,bye bye"

	// 使用 strings.Index() 函数在字符串中搜索另外一个子串
	comma := strings.Index(example, ",")
	fmt.Println(comma)  // 返回字节的下标

	pos := strings.Index(example[comma:], "bye")
	fmt.Println(example[comma:], "----", pos)

	fmt.Println(example[comma+pos:]) // 查找bye


}
