package main

//import "fmt"
//
//func main() {
//	a := []string{"北京", "上海", "深圳"}
//	for i := 0; i < len(a); i++ {
//		fmt.Println(a[i])
//	}
//
//	for index, val := range a {
//		fmt.Println(index, "--", val)
//	}
//
//	// 基于数组定义切片
//	b := [5]int{55, 56, 57, 58, 59}
//	c := b[1:4]
//	fmt.Println(b)
//	fmt.Println(c)
//
//	// 容量
//	s := []int{2, 3, 5, 7, 11, 13}
//	fmt.Println(s)
//	fmt.Printf("长度:%v 容量 %v\n", len(s), cap(s))
//
//	d := s[:2]
//	fmt.Println(d)
//	fmt.Printf("长度:%v 容量 %v\n", len(d), cap(d))
//
//	e := s[1:3]
//	fmt.Println(e)
//	fmt.Printf("长度:%v 容量 %v\n", len(e), cap(e))
//
//	// 使用make函数构造切片
//	f := make([]int, 2, 10)
//	fmt.Println(f)
//	fmt.Println(len(f))
//	fmt.Println(cap(f))
//
//	// append添加元素和切片扩容
//	var numSlice []int
//	for i := 0; i < 10; i++ {
//		numSlice = append(numSlice, i)
//		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
//	}
//
//	var citySlice []string
//	// 追加一个元素
//	citySlice = append(citySlice, "北京")
//	// 追加多个元素
//	citySlice = append(citySlice, "上海", "广州", "深圳")
//	// 追加切片
//	g := []string{"成都", "重庆"}
//	citySlice = append(citySlice, g...)
//
//	fmt.Println(citySlice)
//
//	s1 := []int{100, 200, 300}
//	s2 := []int{400, 500, 600}
//	s3 := append(s1, s2...)
//	fmt.Println(s3)
//
//	h := []int{1, 2, 3, 4, 5}
//	j := []int{6}
//	copy(j, h)
//	fmt.Println(h)
//	fmt.Println(j)
//
//	j[0] = 1000
//	fmt.Println(h)
//	fmt.Println(j)
//
//	// 从切片中删除元素
//	k := []int{30, 31, 32, 33, 34, 35, 36, 37}
//	k = append(k[:2], k[3:]...)
//	fmt.Println(k)
//}
