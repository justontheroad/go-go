package main

import "fmt"

type human struct {
	Sex int
}

//Student 学生，定义结构体
type Student struct {
	//Name 名称
	Name string
	//Score 成绩
	Score float32
}

//Teacher 教师
type Teacher struct {
	string
	int
}

//NewStudent 学生
type NewStudent struct {
	//human嵌套结构体，嵌入结构的名称充当字段名称
	human
	//Name 名称
	Name string
	//Score 成绩
	Score float32
	//匿名结构体
	Contact struct {
		Phone int
	}
}

func main() {
	//声明结构体
	var mimi Student = Student{}
	mimo := Student{}
	fmt.Println(mimi, mimo)
	//初始化结构体
	wuwu := &Student{Name: "wuwu", Score: 100}
	var lili Student = Student{"lili", 100}
	var lilei Student = Student{Name: "lili", Score: 100}
	//使用new初始化结构体
	var dudu = new(Student)
	//访问结构体成员
	mimi.Name = "mimi"
	mimi.Score = 100
	mimo.Name = "mimo"
	mimo.Score = 100
	dudu.Name = "dudu"
	dudu.Score = 99.5
	//匿名strcut
	huhu := &struct {
		Name  string
		Score float32
	}{Name: "huhu", Score: 99.5}

	missLi := &Teacher{"li mei", 30}
	fmt.Println(wuwu, lili, lilei, dudu, huhu, missLi)
	//结构体作为函数参数
	printStudent(mimi)
	//结构体指针作为函数参数
	pringPStudent(&mimi)
	//结构体比较
	fmt.Println(mimi == mimo)

	newS := NewStudent{human: human{Sex: 1}, Name: "kaka", Score: 99} //嵌入结构的名称充当字段名称
	// newS := NewStudent{Sex: 1, Name: "kaka", Score: 99} //错误方式1，cannot use promoted field human.Sex in struct literal of type NewStudent
	// newS := NewStudent{human{Sex: 1}, Name: "kaka", Score: 99} //错误方式2，mixture of field:value and value initializers
	//嵌套的匿名结构体，无法使用字面量初始化
	newS.Contact.Phone = 13000000000
	// newS.Phone = 13008800000 //嵌套的匿名结构体，成员变量无法直接访问
	newS.human.Sex = 1
	newS.Sex = 2
	fmt.Println(newS)
}

func printStudent(student Student) {
	fmt.Printf("Student name: %s, score: %f\n", student.Name, student.Score)
}

func pringPStudent(student *Student) {
	fmt.Printf("Student name: %s, score: %f\n", student.Name, student.Score)
}
