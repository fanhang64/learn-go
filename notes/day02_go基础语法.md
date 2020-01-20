# Go数据类型

### 流程控制switch

使用`switch case`可以对多个值进行判断，简化繁琐的if语句

```go
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

// 可以在switch语句中初始化n的值
switch n:=3;n {
	case 1:
		fmt.Println("111")		
	case 3:
		fmt.Println("333")
	default:
		fmt.Println("moren")
}
```

一个分支可以有多个值，多个case值中间使用英文逗号分隔。

```go
switch n:=3;n {
	case 1,3,5,7,9:
		fmt.Println("奇数")
	case 2,4,6,8,0:
		fmt.Println("偶数")
}
```

分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量（此时switch没有表达式，如果添加表达式case 条件会报error）

```go
age := 19
switch {  // 此时switch没有表达式，如果添加表达式case 条件会报error
	case age > 18:
			fmt.Println("大于18")
	case age < 18:
			fmt.Println("小于18")
	case age == 18:
			fmt.Println("等于18")
}
```

兼容c语言穿透，C语言默认需要加break防止穿透。go相反不需要加，如果想同C语言一样，可以使用`fallthrough`。

```go
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
```



**注意：**

1. 一个switch只能有一个`default`分支

   

### goto(跳转到指定标签)

`goto`语句通过标签进行代码间的无条件跳转。

**实例：**

```go
package main

import "fmt"

// 一般使用
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
```

### break(跳出循环)和continue(跳过本次循环)

`break`语句可以结束`for`、`switch`和`select`的代码块。

```go
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
```

**注意：**

- `continue`语句可以结束当前循环，仅限在`for`循环内使用。
- `break`语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的`for`、`switch`和 `select`的代码块上。



