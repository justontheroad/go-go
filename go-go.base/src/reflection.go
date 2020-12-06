package main

import (
	"fmt"
	"reflect"
)

//User，用户信息
type User struct {
	UID  int
	Name string
	Age  int
	Sex  int
}

//Name，获取用户姓名
func (u User) Info() {
	fmt.Println(u.UID, u.Name, u.Age, u.Sex)
}

type Vip struct {
	User
	Level int
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("type of", t.Name(), t.Kind())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("non-struct type")
		return
	}

	v := reflect.ValueOf(o)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func (u User) Hello(name string) {
	fmt.Printf("Hello %s, my name is %s\n", name, u.Name)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("can't set")
		return
	}

	v = v.Elem()
	f := v.FieldByName("Name")

	if !f.IsValid() {
		fmt.Println("bad!")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("TEST")
	}
}

func main() {
	u := User{1, "tmc", 20, 1}
	u.Info()
	Info(u)
	// Info(&u)

	v := Vip{User: User{1, "tmc", 20, 1}, Level: 1}
	t := reflect.TypeOf(v)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))

	// 通过反射修改对象状态
	i := 123
	ti := reflect.TypeOf(&i)
	ti = ti.Elem()
	fmt.Printf("element name: '%v', element kind: '%v'\n", ti.Name(), ti.Kind())
	ri := reflect.ValueOf(&i)
	ri = ri.Elem()
	ri.SetInt(456)
	// ri.Elem().SetString("456") // call of reflect.Value.SetString on int Value
	fmt.Println(i)
	Set(&u)
	fmt.Println(u)

	// 调用方法
	ru := reflect.ValueOf(u)
	invoke := ru.MethodByName("Hello")
	// 构造函数参数
	args := []reflect.Value{reflect.ValueOf("lilei")}
	// 反射调用函数
	invoke.Call(args)

	var a int
	// 取变量a的反射类型对象
	typeOfA := reflect.TypeOf(a)
	// 根据反射类型对象创建类型实例
	aIns := reflect.New(typeOfA)
	// 输出Value的类型和种类
	fmt.Println(aIns.Type(), aIns.Kind())
}
