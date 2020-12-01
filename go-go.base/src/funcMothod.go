package main

import "fmt"

//Animal动物
type Animal struct {
	Name   string
	weight float32
}

//SI,定义SI类型的string
type SI string

//Receiver 值类型
func (a Animal) call() {
	fmt.Printf("Animal %s call \n", a.Name)
}

//不存在方法重载
// func (a Animal) call(i int) {
// 	fmt.Printf("Animal %s call %d \n", a.Name, i)
// }

//Receiver 指针类型
func (a *Animal) callExchange() {
	//Receiver可能为nil
	if a == nil {
		fmt.Println("Animal is nil")
		return
	}
	a.Name = "dog"
	fmt.Printf("Animal %s call \n", a.Name)
}

func (a *Animal) callWeight(weight float32) {
	a.callExchange()
	//method内部可以访问指针的私有成员 —— 同一个package内，所有私有成员都是可见的
	a.weight = weight
	fmt.Printf("Animal weight:%f \n", a.weight)
}

func (si *SI) print() {
	fmt.Println(*si, si)
}

func main() {
	a := &Animal{Name: "cat"}
	a.call()
	a.callExchange()
	var b *Animal
	b.callExchange()

	// 可以使用值或指针来调用方法，编译器会自动完成转换
	c := Animal{Name: "lion"}
	d := &Animal{Name: "lion"}
	c.call()
	d.call()
	// 方法是函数的语法糖，因为receiver其实就是
	// 方法所接收的第1个参数（Method Value vs. Method Expression）
	var si SI = "test"
	si.print()
	(*SI).print(&si)

	w := &Animal{Name: "shnak", weight: 10}
	w.callWeight(20.00)

	hello("world")
	fmt.Println(increment(1, 2, 3, 4, 5))
	i, j := swap(2, 1)
	fmt.Println(i, j)
	x, y := exchange("x", "y")
	fmt.Println(x, y)
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	//匿名函数
	fn := func(word string) {
		fmt.Printf("hello %s\n", word)
	}
	fn("world")
	add := closureAdd(10)
	fmt.Println(add(20))
	fmt.Println(add(30))
	fmt.Println(add(40))

	//defer，逆序调用
	defer fmt.Print("A")
	defer fmt.Print("B")
	defer fmt.Print("C")
	//defer与匿名函数
	for i := 0; i < 2; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	debug()
}

//hello 普通函数
func hello(word string) {
	fmt.Printf("Hello %s\n", word)
}

//increment 不定长变参
func increment(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

//swap，多返回值
func swap(x, y int) (int, int) {
	return y, x
}

//exchange，命名返回值参数
func exchange(x, y string) (a, b string) {
	a, b = y, x // a, b，函数执行前，a, b已定义
	// a, b := y, x // no new variables on left side of :=
	return
}

//getSequence函数，返回一个闭包函数
func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//closureAdd函数，返回一个闭包函数
func closureAdd(x int) func(int) int {
	fmt.Printf("%p\n", x)
	return func(y int) int {
		fmt.Printf("%p\n", x)
		return x + y
	}
}

func debug() {
	// panic("Panic in debug") //panic先于defer执行，defer不会正常注册
	//defer 用于recover恢复
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in debug")
		}
	}()
	//recover只有在defer调用的函数中有效
	// if err := recover(); err != nil {
	// 	fmt.Println("Recover in debug")
	// }
	panic("Panic in debug")
}
