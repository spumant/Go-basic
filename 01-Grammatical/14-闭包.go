package main

//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//	/*
//		变量 f 是一个函数并且它引用了其外部作用域中的 x 变量，此时 f 就是一个闭包。 在 f 的生
//		命周期内，变量 x 也一直有效。
//	*/
//
//	f := adder()
//	fmt.Println(f(10))
//	fmt.Println(f(20))
//	fmt.Println(f(30))
//
//	f1 := adder()
//	fmt.Println(f1(40))
//	fmt.Println(f1(50))
//
//	jpgFunc := makeSuffixFunc(".jpg")
//	txtFunc := makeSuffixFunc(".txt")
//	fmt.Println(jpgFunc("test"))
//	fmt.Println(txtFunc("test"))
//
//}
//
//func adder() func(int) int {
//	var x int
//	return func(y int) int {
//		x += y
//		return x
//	}
//}
//
//func makeSuffixFunc(suffix string) func(string) string {
//	return func(name string) string {
//		if !strings.HasSuffix(name, suffix) {
//			return name + suffix
//		}
//		return name
//	}
//}
