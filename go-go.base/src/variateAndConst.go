package main

import (
	"fmt"
	"strconv"
)

// 全局变量
var (
	aa     int    = 11
	bb, ss string = "he", "llo"
)

// 定义单个常量
const MIN int = 1

// 省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型
const MAX = 99

// 定义多个常量
const TEST1, TEST2 = 1, 2
const (
	TEXT1 = "A"
	TEXT2 = "B"
	TEXT3 = len(TEXT1)
)

const (
	VAL1 = iota // 0
	VAL2        // 1
	VAL3        // 2
)

// 星期枚举
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	var a int     // 变量声明
	a = 1         // 变量赋值
	var b int = 2 // 变量声明并赋值
	var c = 3     // 省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型
	d := 4        // 简写格式
	fmt.Println(a, b, c, d)
	fmt.Println(aa, bb, ss)

	var (
		z, x = 1, 2
	)
	fmt.Println(z, x)

	// 类型转换
	var f = 5.0
	ii := int(f)
	// bb := bool(ii) // cannot convert ii (type int) to type bool
	fmt.Println(f, ii, bb)
	// 字符串转换
	var ia = 66
	sa := string(ia)       // 转换为B
	ss := strconv.Itoa(ia) // 转换为数字字符串
	var ba = false
	sb := strconv.FormatBool(ba)
	fmt.Println(ia, sa, ss, ba, sb)

	fmt.Println(MIN, MAX)
	fmt.Println(TEST1, TEST2)
	fmt.Println(TEXT1, TEXT2, TEXT3)
	fmt.Println(VAL1, VAL2, VAL3)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
