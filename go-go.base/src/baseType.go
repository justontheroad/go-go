package main

import (
	"fmt"
	// "math/cmplx"
)

func main() {
	// 整形
	var i8 int8 = -128           // -128~127
	var i16 int16 = -32768       // -32768~32767
	var i32 int32 = -2147483648  // -2147483648~2147483647
	var i64 int64 = -2147483648  // -2^64/2~2^64/2-1
	var ui8 uint8 = 255          // 0~255
	var ui16 uint16 = 65535      // 0~65535
	var ui32 uint32 = 4294967295 // 0~2147483649
	var ui64 uint64 = 4294967295 // 0~2^64-1
	fmt.Println(i8, i16, i32, i64, ui8, ui16, ui32, ui64)

	// 浮点形
	var f32 float32 = 0.01
	var f64 float64 = 0.001
	fmt.Println(f32, f64)

	// 复数
	var c64 complex64 = complex(1, 2)   // 1+2i
	var c128 complex128 = complex(3, 4) // 3+4i
	fmt.Println(c64, c128)

	// 布尔
	var bl bool = true
	fmt.Println(bl)

	// byte与rune
	var chA byte = 65
	var chB byte = '\x41' // 在 ASCII 码表中，A 的值是 65，使用 16 进制表示则为 41
	var chU int = '\u0041'
	fmt.Printf("%d - %d - %d\n", chA, chB, chU) // integer
	fmt.Printf("%c - %c - %c\n", chA, chB, chU) // character
	fmt.Printf("%X - %X - %X\n", chA, chB, chU) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", chA, chB, chU)   // UTF-8 code point

	// 字符串
	var str string = "hello"
	fmt.Println(str)
	var str1 string       // 声明一个字符串变量
	str1 = "Hello world"  // 字符串赋值
	ch := str1[0]         // 取字符串的第一个字符
	str2 := "Hello world" // 直接初始化，推导为string类型
	// str2[0] = 'a'           // 编译错误，不支持初始化后修改内容
	str1 = str1 + str2 // str1的值变为了Hello worldHello world
	fmt.Println(ch, str1)

	// 数组
	var arr [10]int
	arr[2] = 1
	fmt.Println(arr)
	// arr2 := [10]int{1,2,3,4,5,6,7,8,9}
	// fmt.Println(arr2)
	// arr3 := [...]int{1,2}
	// fmt.Println(arr3)
	// 切片
	s1 := arr[1]
	s2 := arr[2:5] // arr[2,3,4]
	fmt.Println(s1, s2)
	// 字典
	var kv map[string]string
	kv = map[string]string{"aa": "test"}
	fmt.Println(kv)
	kv1 := make(map[string]string)
	fmt.Println(kv)
}
