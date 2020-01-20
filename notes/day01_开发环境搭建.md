## GO语言开发环境搭建



GO语言官方镜像站：https://golang.google.cn/

```
> go version  查看go版本号
go version go1.12.6 linux/amd64

```

### Linux下安装

我们在版本选择页面选择并下载好`go1.11.5.linux-amd64.tar.gz`文件：

```bash
wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
```

将下载好的文件解压到`/usr/local`目录下：

```bash
mkdir -p /usr/local/go  # 创建目录
tar -C /usr/local/go zxvf go1.11.5.linux-amd64.tar.gz. # 解压
```

如果提示没有权限，加上`sudo`以root用户的身份再运行。执行完就可以在`/usr/local/`下看到go目录了。

配置环境变量： Linux下有两个文件可以配置环境变量，其中`/etc/profile`是对所有用户生效的；`$HOME/.profile`是对当前用户生效的，根据自己的情况自行选择一个文件打开，添加如下两行代码，保存退出。

```bash
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

修改`/etc/profile`后要重启生效，修改`$HOME/.profile`后使用source命令加载`$HOME/.profile`文件即可生效。 检查：

```bash
~ go version
go version go1.11.5 linux/amd64
```

### **配置GOPATH**

`gopath`环境变量, 用来表明go项目存放的路径(工作目录)

`goroot`是安装go文件位置

```python
# 1) 
添加系统环境变量GOPATH 目录为D:/GO
    
# 2) 新建3个目录
bin/  src/ pkg/

# 3) 将D:/go/bin添加到path环境变量

# 4) 命令 go env 查看go配置环境变量
```



### **go项目结构**

个人开发

```
bin/    存放编译后二进制文件

pkg/    粗放编译后库文件

src/    存放源码文件
  项目1/
  项目2/
```

现在流行的

```
bin/    存放编译后二进制文件

pkg/    粗放编译后库文件

src/    存放源码文件
  github.com/
      fahang64/
  		项目1/
  		项目2/
      abc/      # 作者名字
```

企业

```
bin/    存放编译后二进制文件

pkg/    粗放编译后库文件

src/    存放源码文件
   github.com/
   golang.org/
   code.xxcompany.com/
        front/
        backend/
            项目1/
            项目2/
        others/
```

**实例：**

举个例子：张三和李四都有一个名叫`learn-go`的项目，那么这两个包的路径就会是：

```go
import "github.com/zs/lean-go"
```

和

```go
import "github.com/ls/lean-go"
```

以后我们从github上下载别人包的时候，如：

```bash
go get github.com/jmoiron/sqlx
```

那么，这个包会下载到我们本地`GOPATH`目录下的`src/github.com/jmoiron/sqlx`。

#### **第一个go程序**

```go
// src/xxx.com/day01/01helloworld/main.go

package main

import "fmt"

func main()  {
	fmt.Println("Hello world")
}
```

#### **执行编译**

`go build`

```go
// 在项目路径下

// 法一： 

> cd src/xxx.com/day01/01helloworld/
> go build
> ./01helloworld

// 法二:

> cd ~
> go build xxx.com/day01/01helloworld/   // 从src后开始写，写到项目目录
> ./01helloworld   // 在home目录生成执行文件。从哪里go build执行，生成执行文件会在哪里

// 法三:  修改名字
> cd src/xxx.com/day01/01helloworld/
> go build  -o  hello.exe 
> ./hello

// 法四:  
> cd src/xxx.com/day01/01helloworld/
> go run main.go    // 直接执行文件 类似脚本

// 法五:
> cd src/xxx.com/day01/01helloworld/
> go install    // 可以理解为分为两部: 1) 先编译得到一个可执行文件； 2）讲可执行文件拷贝到GOPATH/bin 中

```

#### **交叉编译**

go支持跨平台编译

例如: windows下平台下编译在linux平台执行的可执行文件

默认我们`go build`的可执行文件都是当前操作系统可执行的文件，如果我想在windows下编译一个linux下可执行文件，那需要怎么做呢？

只需要指定目标操作系统的平台和处理器架构即可：

```bash
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

*使用了cgo的代码是不支持跨平台编译的*

然后再执行`go build`命令，得到的就是能够在Linux平台运行的可执行文件了。

Mac 下编译 Linux 和 Windows平台 64位 可执行程序：

```bash
> CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build   # 直接执行
> CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

Linux 下编译 Mac 和 Windows 平台64位可执行程序：

```bash
> CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build   # 直接执行
> CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

Windows下编译Mac平台64位可执行程序：

```bash
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build
```

#### go语言文件基本结构

```go
package main   // main包生成执行文件， 其他名字作为模块

import "fmt"   // 导入语句

// 函数外只能放置标识符（变量，常量，函数，类型）的声明
// fmt.Println("xxxxx")  // error 只能放到函数里

// 程序的入口函数
func main()  {
	fmt.Println("Hello world")
}
```

### 变量

go的标识符(程序员定义的)只能由数字，字母，下划线组成

**变量**

Go语言的变量声明格式为：

```go
var 变量名 变量类型
```

变量声明以关键字`var`开头，变量类型放在变量的后面，行尾无需分号。 

例如：

```go
var name string
var age int
var isOk bool
```

#### 批量声明

每声明一个变量就需要写`var`关键字会比较繁琐，go语言中还支持批量变量声明：

```go
var (
    a string
    b int
    c bool
    d float32
)
```

#### 变量的初始化

Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。

每个变量会被初始化成其类型的默认值，例如： 整型和浮点型变量的默认值为`0`。 字符串变量的默认值为`空字符串`。 布尔型变量默认为`false`。 切片、函数、指针变量的默认为`nil`。

当然我们也可在声明变量的时候为其指定初始值。变量初始化的标准格式如下：

```go
var 变量名 类型 = 表达式
```

例如：

```go
var name string = "fanhang64"
var age int = 20
```

或者一次初始化多个变量

```go
var name, age = "fanhang64", 20
```

#### 类型推导

有时候我们会将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。

```go
var name = "fanhang64"
var age = 20
```

#### 短变量声明

在函数内部，可以使用更简略的 `:=` 方式声明并初始化变量。

```go
package main

import (
	"fmt"
)
// 全局变量m
var m = 100

func main() {
	n := 10
	m := 200 // 此处声明局部变量m
	fmt.Println(m, n)
}
```

#### 匿名变量

在使用多重赋值时，如果想要忽略某个值，可以使用`匿名变量`

匿名变量用一个下划线`_`表示，

例如：

```go
func foo() (int, string) {
	return 20, "fanhang64"
}
func main() {
	x, _ := foo()  // 忽略"fanhang64"
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
```

**注意：**

1. 函数外的每个语句都必须以关键字开始（var、const、func等）
2. `:=`不能在函数外使用。
3. `_`多用于占位，表示忽略值。
4. 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。 

**实例：**

```go
package main

import "fmt"

// GO中使用驼峰命名法

// 声明变量
// var name string
// var age int
// var isOk bool

// 批量声明
var (
	name string   // 默认值 ""
	age int       // 默认值 0
	isOk bool     // 默认值 false
	gender bool
    score float32
)

func main()  {
	name = "fanhang64"
	age = 16 // 变量赋值
	isOk = true 

	// go语言中函数内的变量声明必须使用，不使用编译不过去
	// 全局变量可以声明不使用
	// var haha string  // 声明 不使用error
	//  gender 全局变量声明  不使用ok

	fmt.Println(age)  // 打印指定变量，结尾添加换行
	fmt.Printf("name:%s\n", name)  // 占位符
	fmt.Print(isOk)  // 输出isok内容在终端中
	fmt.Println()  // 空行

	var s1 string = "fhz"  // 声明同时赋值
	fmt.Println(s1)

	// 类型推导（根据值判断类型）
	var s2 = "20"
	fmt.Println(s2)  // ok 

	// 短变量声明
	// 只能在函数里用
	s3 := "hahah"
	fmt.Println(s3)

	// s3 := "error "  // 同一个作用域{}中不能重复声明变量

	// 匿名变量  _ 可以忽略返回值，不占用内存，可以重复声明
}
```

### 常量

常量在定义的时候必须赋值，定义之后在程序的运行期间不会在改变。

```
const pi = 3.14
const e = 2.7182
```

多个变量一起声明：

```go
// const 常量
// 定义了常量之后不能修改，在程序运行期间不会改变
const pi = 3.1415926

// 批量声明常量
const (
	STATUSOK = 200
	NOTFOUND = 404
)

```

批量声明常量值，如果某一行声明后，没有赋值，同上一行

```go
// n2 n3 默认和n1相同
const (
	n1 = 100   //  100
	n2         // 100
	n3         // 100

)

const (
	m1 = 100   // 100
	m2         // 100
	m3         // 100
	m4 = 200   // 200
	m5         // 200

)
```

#### iota常量计数器

`iota`是go语言的常量计数器，只能在常量的表达式中使用。

```go
const (
	a1 = iota  // 0 
	a2         // 1  相当于 a2 = iota
	a3         // 2  相当于 a3 = iota
	a4 = iota  // 3  
)

```

**注意：**

1. `iota`只能在常量的表达式中使用
2. 在声明常量时，第一次出现时为0
3. const中每新增一行常量声明将使iota计数加1

**实例：**

```go
package main

import "fmt"


// const 常量
// 定义了常量之后不能修改，在程序运行期间不会改变
const pi = 3.1415926

// 批量声明常量
const (
	STATUSOK = 200
	NOTFOUND = 404
)

// 批量声明常量值，如果某一行声明后，没有赋值，同上一行
// n2 n3 默认和n1相同
const (
	n1 = 100
	n2
	n3

)

const (
	m1 = 100
	m2
	m3
	m4 = 200
	m5

)


// iota 是go语言中常量计数器, 只能在常量的表达式中使用
// 在const中 第一次出现被重置为0
// const中每新增一行常量声明将使iota计数一次
const (
	a1 = iota  // 0 
	a2         // 1  相当于 a2 = iota
	a3         // 2  相当于 a3 = iota
	a4 = iota  // 3  
)


// iota常见示例：
// 1） 使用_跳过某些值
const (
	b1 = iota  // 0
	b2         // 1
	_          // 忽略掉2
	b3         // 3
)

// 2）插队
const (
	c1 = iota   // 0
	c2 = 100    // 100   const中每新增一行常量声明将使iota计数一次  相当于1的位置
	c3 = iota   // 2
	c4          // 3   == c4 = iota
)

// 3） 多个iota声明在一行
const (
	d1, d2 = iota+1, iota +2   // 0+1 0+2
	d3, d4 = iota +1, iota +2  // 1+1 1+2
	d5, d6                     // 2+1 2+2 同上 d5,d6 = = iota +1, iota +2 
)

// 4） 定义数量级

const (
	_ = iota
	KB = 1 << (10 * iota)  // 1左移10位即1024
	MB = 1 << (10 * iota)  // 1左移10 *2 位
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
func main()  {

	// pi = 123  // 修改常量 error
	fmt.Println(n1, n2, n3)  // 100 100 100
	fmt.Println(m1, m2, m3, m4, m5)  // 100 100 100  200 200
	iota := 1000   // ok 
	
	fmt.Println(iota)  // 1000

	fmt.Println(a1, a2, a3, a4)  // 0 1 2 3
	fmt.Println(b1, b2, b3)  // 0 1 3

	fmt.Println(c1, c2, c3, c4)  // 0 100 2 3
	fmt.Println(d1, d2, d3, d4)  // 1 2 2 3

	fmt.Println(KB, MB, GB, TB, PB)
	
}
```

### 基本数据类型

Go语言中有丰富的数据类型，除了基本的整型、浮点型、布尔型、字符串外，还有数组、切片、结构体、函数、map、通道（channel）等

#### 整型

整型分为以下两个大类： 

按长度分为：int8、int16、int32、int64 

对应的无符号整型：uint8、uint16、uint32、uint64

**注意：** 

1. `uint8`就是我们熟知的`byte`型，`int16`对应C语言中的`short`型，`int64`对应C语言中的`long`型。
2. 在使用`int`和 `uint`类型时，不能假定它是32位或64位的整型，而是考虑`int`和`uint`可能在不同平台上的差异。
3. 获取对象的长度的内建`len()`函数返回的长度可以根据不同平台的字节长度进行变化。

**实例：**

```go
package main

import "fmt"

// 整型
func main()  {
	var a = 10  // 类型推导

	fmt.Printf("%d\n", a)

	// 八进制
	b := 077  
	fmt.Printf("%d %o\n", b, b)

	// 十六进制数
	c := 0xff
	fmt.Printf("%d, %x\n", c, c)
	
	// 查看变量的类型
	fmt.Printf("%T\n", c)  // 默认int类型

	// 声明一个int8类型的变量
	d := int8(9)
	fmt.Printf("%T\n", d)  // 默认int8类型

}
```

#### 浮点型

Go语言支持两种浮点型数类型：`float32`和`float64`。

`float32`最大值的常量定义：`math.MaxFloat32`。 

`float64` 最大值的常量定义：`math.MaxFloat64`。

**注意：**

- 两种类型不能相互赋值

**实例：**

```go
package main

import "fmt"
// import "math"

func main()  {
	// math.MaxFloat32  // 32位浮点数最大值
	// math.MaxFloat64  // 64位浮点数最大值

	f1 := 1.2345
	fmt.Printf("%T\n", f1)  // 默认go语言中浮点数是float64

	f2 := float32(1.23455)
	fmt.Printf("%T\n", f2)  // 显示声明float32
	
	// f1 = f2  // error 不能直接把float64传给float32
}
```

#### 复数

两种类型：`complex64`和`complex128`

```go
var c1 complex64
c1 = 1 + 2i
var c2 complex128
c2 = 2 + 3i

fmt.Println(c1)
fmt.Println(c2)
```

**注意：**

- 复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。

#### 布尔类型

Go语言中以`bool`类型进行声明布尔型数据，布尔型数据只有`true（真）`和`false（假）`两个值。

**注意：**

1. 布尔类型默认值为false
2. 不允许将整型强制转换为bool类型
3. bool类型无法进行数值运算，也无法与其他类型进行转换

**实例：**

```go
package main

import "fmt"

func main(){
	b1 := true

	var b2 bool  // 默认false

	fmt.Println(b1)
	fmt.Printf("%T value: %v", b2, b2)  // %v 打印变量的值 无论任何类型

	b3 := 10  
	
	// b4 := bool(b3) // error 不能将整型转为布尔

	// b5 := b2 + b3  // error 不能进行数值运算
	
	fmt.Println(b3)
}
```

**fmt包总结**

**实例：**

```go
package main

import "fmt"

func main()  {
	var n = 100

	// fmt包总结
	fmt.Printf("%T\n", n)  // 查看类型
	fmt.Printf("%v\n", n)  // 打印变量的值
	fmt.Printf("%b\n", n)   
	fmt.Printf("%o\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%x\n", n)

	// 字符串
	var s = "hello"

	fmt.Printf("%s\n", s)  // hello
	fmt.Printf("%v\n", s)  // hello
	fmt.Printf("%#v\n", s)  // "hello"  打印字符串
}
```



#### 字符串

在Go语言中使用`""`双引号包裹，内部实现使用`UTF8`编码。

使用单引号`''`包裹的是一个字符。

**例如：** 

```go
s1 := "hello"
s2 := "你好"

// 使用【反引号】定义多行字符串
s3 = `
定义多行字符串。。。。。内容
`

// 字符串拼接
ss := "123" +
    "asdasdasd" +
    "hello world"

fmt.Println(ss)
```

**常用函数**

| 函数                                                         | 功能                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| len(str)                                                     | 判断字符串占用字节长度。一个字母占用1个字节，汉字3个字节。返回int |
| strings.Sprintf(”格式“, 变量名)                              | 拼接字符串返回string类型。                                   |
| strings.Split("字符串", "使用分割的子串")                    | 根据子串将字符串进行分割。                                   |
| strings.Contains(字符串, 子串)                               | 判断字符串是否包含子串，返回bool类型，如果 substr 为空，则返回 true |
| strings.HasPrefix(字符串, 子串)  strings.HasSuffix(字符串, 子串) | 判断字符串是否是以【子串】作为前缀或后缀。返回bool           |
| strings.Index(字符串, 子串)  strings.LastIndex()             | 返回第一次或最后一次出现子串的下标，下标从0开始，返回int     |
| strings.Join(a [] strings,  "拼接子字符串")                  | 使用子串拼接字符串数组                                       |
| strings.ContainsRune(s string, r rune) bool                  | 判断字符串s中是否包含字符r                                   |

**实例：**

```go
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
```

###### byte和rune类型

字符串由每一个字符组成，可以通过遍历或者索引下标的方式获取每个元素。

 字符用单引号`''`包裹，默认类型为`rune`类型，如果要定义`byte`类型变量需要手动定义。

例如：

```go
c1 := "哈"
c2 := '哈'  // 默认rune类型 也是(int32) 类型
fmt.Printf("c1: %T, c2: %T\n", c1, c2)  // c1: string, c2: int32  

c3 := "H"
c4 := 'H'  // 默认是rune类型，也是int32
c5 := byte('H')  // 强制转换为byte类型，也是uint8类型
fmt.Printf("c3: %T, c4:%T, c5:%T\n", c3, c4, c5)  // c3: string, c4:int32
```

**注意：**

1. `byte`类型又叫`uint8`类型，代表`ascii`码的一个字符。
2. `rune`类型，代表一个`utf-8`字符，当需要处理中文，日文或其他字符时，需要使用`rune`类型，该类型实际是int32。

**实例：**

```go
// 字符串的遍历
func main() {
	s := "hello世界"
	for i := 0; i < len(s); i++ {  // 按照byte遍历字符串
		fmt.Printf("%v(%c) ", s[i], s[i])   // 打印字符, 可以打印出,中文出现乱码
	}
	fmt.Println()

	for _, r := range s {         // 按照rune类型遍历字符串
		fmt.Printf("%v(%c) ", r, r)  // ok
	}
	fmt.Println()
}

// 字符串的修改
s2 := "哈哈哒"  // 表示 '哈' '哈' '哒' 三个字符
// s2[0] = "吧"  // error 字符串不能修改，需要强制转换为切片

s3 := []rune(s2)  // 将s2强制转为rune切片

// s3[0] = "把"  // error 
s3[0] = '把'  // ok [哈 哈 哒] s[0]= '' 应该是字符
fmt.Println(string(s3))  // 将切片转为字符串

// 类型转换
// 注意： bool不能和其他类型转换

// 整数转浮点数
m := 10
var f float64
f = float64(m)
fmt.Printf("%T, %f\n", f, f)  // float64, 10.000000

// 浮点数转整数  
// 小数部分被丢弃
var v1 float32 = 0.999999
fmt.Println(int(v1))  // 0
v1 = -0.999999
fmt.Println(int(v1))  // 0
```

### 流程控制

#### `if`条件判断:

go语言固定与`if`匹配的左括号必须与`if和表达式`放在一行，`{`放在其他位置会出发编译错误。同理，`else`，`else if`也要和`{`放在一行。

格式：

```go
if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
```

**if条件判断特殊用法**

实例：

```go
if age := 18; age > 18{
    fmt.Println("大于18了")
}else{
    fmt.Println("小于18")
}
```

**实例：**

```go
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
    if 1 {  // error 条件必须为bool类型
        
    }
}
```



#### for循环

**格式：**

```go
for 初始语句;条件表达式;结束语句{
    循环体语句
}
```
**注意：**

- `for`初始语句可以被忽略，但是初始语句后的分号必须要写
- `for`循环的初始语句和结束语句都可以省略
- `for`循环可以通过`break`、`goto`、`return`、`panic`语句强制退出循环。



#### for无限循环

```go
for{
    循环体语句
}
```

#### for range循环

Go语言中可以使用`for range`遍历数组、切片、字符串、map 及通道（channel）。

 通过`for range`遍历的返回值有以下规律：

1. 数组、切片、字符串返回索引和值。
2. map返回键和值。
3. 通道（channel）只返回通道内的值。



**实例：**

```go
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
	// 同变种2, 两个分号忽略了
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
	for i := range "helloworld"{
		fmt.Println(i)  // 索引从0开始
	}
   
	s := "hello世界"
	for i,v := range s{  // i索引 v字符
		fmt.Println(i, v)
		fmt.Printf("%v(%c)\n", v, v)  
	}
}
```


