package main

import "fmt"

func main() {
	// 声明切片
	var iSlice []int
	fmt.Println(iSlice)
	// 初始化切片
	s := []int{1, 2, 3}
	// 从底层数组获取生成
	ia := []int{10: 1}
	is := ia[1:11]
	fmt.Println(s, ia, is)
	// 使用make创建切片，不指定容量，默认跟长度一致
	mis := make([]int, 2<<2, 2<<3)
	fmt.Println(mis, len(mis), cap(mis))
	// 空切片nil
	var num []int
	if num == nil {
		fmt.Println("切片是空的", len(num), cap(num))
	}

	// Reslice
	as := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n'}
	s1 := as[3:6] // 长度为3，容量为11
	rs := s1[1:3]
	// rs := s1[1:11] // rs 截取 s1时，索引最大值为11(s1的最大容量)，超过时提示索引越界；
	// rs := s1[1:12] // slice bounds out of range [:12] with capacity 11
	fmt.Println(len(as), cap(as), string(s1), len(s1), cap(s1), string(rs), len(rs), cap(rs))

	// Append
	sa := make([]int, 2<<1, 2<<2)
	fmt.Println(len(sa), cap(sa))
	// 长度未超过追加到slice的容量则返回原始slice
	sa = append(sa, 1, 2, 3, 4)
	fmt.Printf("%v, %v, %v, %p\n", len(sa), cap(sa), sa, sa)
	// 超过追加到的slice的容量则将重新分配数组并拷贝原始数据
	sa = append(sa, 1, 2, 3, 4)
	fmt.Printf("%v, %v, %v, %p\n", len(sa), cap(sa), sa, sa)

	// Copy
	cs1 := []int{1, 2, 3, 4, 5}
	cs2 := []int{7, 8, 9}
	copy(cs1, cs2)
	fmt.Println(cs1, len(cs1), cap(cs1))
	cs3 := []int{11, 12, 13, 14, 15}
	copy(cs2, cs3)
	fmt.Println(cs2, len(cs2), cap(cs2))
}
