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
}
