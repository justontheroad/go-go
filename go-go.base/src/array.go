package main

import "fmt"

func main() {
	// 声明数组
	var a [3]int
	// 声明并初始化数组
	var b = [2]int{1, 2}
	// 指定索引初始化数组，指定下标4初始值为1
	var bb = [5]int{4: 1}
	// 省略数组长度
	var bbb = [...]int{1, 1, 1}
	// 简写
	c := [2]int{1, 2}
	// a = b // annot use b (type [2]int) as type [3]int in assignment
	// 访问数组元素
	i := c[1]
	c[0] = 3
	// i := c[2] // invalid array index 2 (out of bounds for 2-element array)
	fmt.Println(a, b, bb, len(bbb), c, i)
	// 数组迭代
	for t := 0; t < len(c); t++ {
		fmt.Println(t, c[t])
	}

	var x = [...]int{9: 1}
	// 指向数组的指针，类型必须和数组一致
	var px *[10]int = &x // var px *([10]int) = &x
	// var px *[100]int = &x // cannot use &x (type *[10]int) as type *[100]int in assignment
	// 指针数组
	var xa, xb, xc = 1, 2, 3
	var ppx = [3]*int{&xa, &xb, &xc} // var ppx = [3](*int){&xa, &xb, &xc}
	fmt.Println(px, ppx)

	// 数组比较，相同类型之间才能比较
	ca, cb, cc := [2]int{1, 2}, [2]int{1, 3}, [2]int{1, 2}
	fmt.Println(ca == cb, ca == cc, ca != cb)

	// 使用new创建数组，返回一个指向数组的指针
	var pn = new([10]int)
	pn[2] = 3
	fmt.Println(pn)

	// 多维数组
	var ma = [3][2]int{
		{1, 2},
		{2, 2},
		{3, 2}}
	fmt.Println(ma)
}
