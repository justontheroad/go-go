package main

import "fmt"

//DataWriter，定义一个数据写入器
type DataWriter interface {
	Write(data interface{}) error
	//嵌套一个另一个接口
	DataWriteLocker
}

//DataWriteLocker，定义数据锁和解锁
type DataWriteLocker interface {
	Lock(data interface{}) error
	Unlock(data interface{}) error
}

//file，定义文件结构，用于实现DataWriter
type file struct {
}

//实现DataWriter接口的Writ方法
func (f file) Write(data interface{}) error {
	//模拟写入
	fmt.Println("Write:", data)
	return nil
}

//实现DataWriteLocker接口的Lock方法
func (f file) Lock(data interface{}) error {
	//模拟锁
	fmt.Println("Lock:", data)
	return nil
}

//实现DataWriteLocker接口的Unlock方法
func (f file) Unlock(data interface{}) error {
	//模拟解锁
	fmt.Println("Unlock:", data)
	return nil
}

//模拟刷洗数据到磁盘
func flush(writer DataWriter) {
	//类型断言
	if f, ok := writer.(file); ok {
		fmt.Println("flush succes", f)
		return
	}

	fmt.Println("flush failed")
}

//type switch
func getType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("the type of i is int")
	case string:
		fmt.Println("the type of i is string")
	case float64:
		fmt.Println("the type of i is float")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	f := new(file)
	// 声明一个DataWriter的接口
	var writer DataWriter
	// 将接口赋值f，也就是*file类型；接口的内部存储的指向这个复制的指针
	writer = f
	fmt.Println(f, writer)
	writer.Lock("data")
	writer.Write("data")
	writer.Unlock("data")
	flush(writer)
	w := file{}
	flush(w)
	// 空接口
	var any interface{}
	any = 1
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)
	any = true
	fmt.Println(any)
	getType(1)
	// 接口转换，向上转型
	var locker DataWriteLocker
	locker = DataWriteLocker(writer)
	locker.Lock("lock data")
	// locker.write("test") // locker.write undefined (type DataWriteLocker has no field or method write)
}
