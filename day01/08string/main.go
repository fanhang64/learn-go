package main

import (
	"fmt"
	"strings"
)

func main()  {

	// 字符串

	path := "\"\\home\\fanzone\\Code\\go\\src\\learn-go\\day01\\08string\\main.go\""

	fmt.Println(path)  // 打印带双引号路径


	// 定义多行的字符串
	s1 := `
	hello world
abcdef13
`
	fmt.Println(s1)

	// 字符串+
	ssa := "123" +
		"asdasdasd" +
		"hello world"

	fmt.Println(ssa)  // ok  123asdasdasdhello world

	// 字符串相关操作
	s3 := "hello world"
	fmt.Println(len(s3))

	// 字符串的拼接
	name := "dream"
	world := "dsb"

	// 法一:
	fmt.Println(name + world)

	// 法二：
	var ss string
	ss = fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss)

	// 字符串分割
	path = "fanzone/code/abc/def"
	ret := strings.Split(path, "/")
	fmt.Println(ret)  // [fanzone code abc def]

	// 包含
	fmt.Println(strings.Contains(path, "code"))  // true
	fmt.Println(strings.Contains(path, "code1"))  // false

	// 前缀 后缀

	fmt.Println(strings.HasPrefix(path, "/"))  // true
	fmt.Println(strings.HasSuffix(path, "def"))  // true

	// 子串出现的位置下标，下标从0开始
	s4 := "abcdecfg"

	fmt.Println(strings.Index(s4, "c"))  // 2 查找第一次出现的位置
	fmt.Println(strings.LastIndex(s4, "c"))  // 5  最后一次出现的位置
	fmt.Println(strings.Index(s4, "a"))  // 0
	fmt.Println(strings.Index(s4, "ABC"))  // 找不到返回-1

	// join操作 拼接操作

	fmt.Println(ret)  // [ fanzone code abc def]

	// 使用+拼接
	fmt.Println(strings.Join(ret, "+"))  //  fanzone+code+abc+def


}