package main

import (
	"fmt"
	"sort"
)

func main() {
	//  声明map
	var mapper map[string]string
	// 初始化map
	mapper = map[string]string{"1": "one", "2": "two"}
	fmt.Println(mapper)
	m1 := map[string]string{"3": "three", "4": "four"}
	// make 创建map
	m2 := make(map[string]string, 2)
	m2["5"] = "five"
	m2["6"] = "six"
	fmt.Println(m1, m2)
	// map迭代
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	// 不需要的键使用_改为匿名变量形式
	for _, v := range m2 {
		fmt.Println(v)
	}
	// 迭代key
	for k := range m2 {
		fmt.Println(k)
	}
	// map元素操作
	coutryCityMap := make(map[string]string, 5)
	// 赋值
	coutryCityMap["France"] = "巴黎"
	coutryCityMap["Italy"] = "罗马"
	coutryCityMap["Japan"] = "东京"
	coutryCityMap["India"] = "新德里"
	coutryCityMap["Test"] = "测试"
	// 删除
	delete(coutryCityMap, "Test")
	// 查看元素是否存在
	t, ok := coutryCityMap["Test"]
	fmt.Println(coutryCityMap)
	if ok {
		fmt.Println(t)
	}
	// 多维map
	mm := make([]map[string]string, 5)
	// v 为mm的拷贝，内部任何修改都不会影响到mm
	for k, v := range mm {
		v = make(map[string]string, 1)
		v["1"] = string(k)
	}
	fmt.Println(mm)
	for k := range mm {
		mm[k] = make(map[string]string, 1)
		mm[k]["1"] = string(k)
	}
	fmt.Println(mm)
	// map排序
	fruitsMap := make(map[string]int, 5)
	fruitsMap["pear"] = 10
	fruitsMap["lemon"] = 20
	fruitsMap["banana"] = 10
	fruitsMap["apple"] = 5
	fruitsMap["orange"] = 2
	// 声明一个切片用于存放map数据
	var sl []string
	// 将map数据遍历复制到切片中
	for k := range fruitsMap {
		sl = append(sl, k)
	}
	sort.Strings(sl)
	fmt.Println(fruitsMap)
}
