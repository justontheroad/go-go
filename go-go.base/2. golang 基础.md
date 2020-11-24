### Go内置关键字（25个均为小写）
break        default           func         interface          select
case         defer             go           map                struct
chan         else              goto         package            switch
const        fallthrough       if           range              type
continue     for               import       return             var 

### 基础数据类型
1. 整形
    1. 8位整型：int8/uint8；长度：1字节，取值范围：-128~127/0~255
    2. 16位整型：int16/uint16；长度：2字节，取值范围：-32768~32767/0~65535
    3. 32位整型：int32（rune）/uint32；长度：4字节，取值范围：-2^32/2~2^32/2-1/0~2^32-1
    4. 64位整型：int64/uint64；长度：8字节，取值范围：-2^64/2~2^64/2-1/0~2^64-1
    5. 
    6. 足够保存指针的 32 位或 64 位整数型：uintptr
2. 浮点型：float32/float64；长度：4/8字节，小数位：精确到7/15小数位
3. 布尔类型：bool
4. 复数：complex64/complex128；长度：8/16字节
5. 字符类型：byte代表UTF-8字符串的单个字节的值，rune，代表单个Unicode字符
    1. byte与rune：byte与rune都属于别名类型。byte是uint8的别名类型，而rune是int32的别名类型
6. 字符串：string
4. 高级数据类型
    1. 数组：array
    2. 切片：slice
    3. 字典：map
    4. 通道：chan
5. 结构体：struct
6. 接口类型：inteface
7. 函数类型：fun

### Go程序结构
Go程序是通过 package 来组织的，只有 package 名称为 main 的包可以包含 main 函数
- 一个可执行程序 有且仅有 一个 main 包 
- 通过 import 关键字来导入其它非 main 包 - 通过 const 关键字来进行常量的定义
- 通过在函数体外部使用 var 关键字来进行全局变量的声明与赋值
- 通过 type 关键字来进行结构(struct)或接口(interface)的声明
- 通过 func 关键字来进行函数的声明

### package
1. 导入 package 的格式
```
import "fmt"
import "strings"
```
```
import (
    "fmt"
    "strings"
)
```
    - 导入包之后，就可以使用格式<PackageName>.<FuncName>
    来对包中的函数进行调用
    - 如果导入包之后 未调用 其中的函数或者类型将会报出编译错误
2. package 别名，当包名接近或相同时，可以使用别名进行区分和调用
```
import std "fmt"
// 使用别名调用
std.Println("hello word")
```
3. 省略调用
```
import . "fmt"
Println("hello word")
```

### 变量的声明与赋值
- 变量的声明格式：var <变量名称> <变量类型>
- 变量的赋值格式：<变量名称> = <表达式> 
- 声明的同时赋值：var <变量名称> [变量类型] = <表达式>
- 变量声明与赋值，简写：<变量名称> = <表达式>
```
var a int
a = 110
var a int = 110
x = 110
```

### 可见性规则
1. 使用 大小写 来决定该 常量、变量、类型、接口、结构或函数 是否可以被外部包所调用：
 - 根据约定，函数名首字母 小写 即为private
 - 函数名首字母 大写 即为public