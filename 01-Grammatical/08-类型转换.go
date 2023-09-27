package main

//import (
//	"fmt"
//	"strconv"
//)
//
//func main() {
//	num1 := 20
//	s1 := strconv.Itoa(num1)
//	fmt.Println(s1)
//
//	num2 := 20.113123
//	/* 参数 1：要转换的值
//	参数 2：格式化类型
//	'f'（-ddd.dddd）、
//	'b'（-ddddp±ddd，指数为二进制）、
//	'e'（-d.dddde±dd，十进制指数）、
//	'E'（-d.ddddE±dd，十进制指数）、
//	'g'（指数很大时用'e'格式，否则'f'格式）、
//	'G'（指数很大时用'E'格式，否则'f'格式）。
//	参数 3: 保留的小数点 -1（不对小数点格式化）
//	参数 4：格式化的类型
//	*/
//	s2 := strconv.FormatFloat(num2, 'f', 2, 64)
//	fmt.Println(s2)
//
//	num3 := 20
//	fmt.Printf("%T\n", num3)
//	/*
//		第二个参数为 进制
//	*/
//	s3 := strconv.FormatInt(int64(num3), 10)
//	fmt.Println(s3)
//
//	s := "1234"
//	i64, _ := strconv.ParseInt(s, 10, 64)
//	fmt.Println(i64)
//
//	str := "3.1415926535"
//	v1, _ := strconv.ParseFloat(str, 32)
//	v2, _ := strconv.ParseFloat(str, 64)
//	fmt.Println(v1, v2)
//}
