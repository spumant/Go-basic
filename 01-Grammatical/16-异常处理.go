package main

//import (
//	"errors"
//	"fmt"
//)
//
//func fn1(x, y int) int {
//	defer func() {
//		err := recover()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}()
//	return x / y
//}
//func readFile(fileName string) error {
//	if fileName == "main.go" {
//		return nil
//	} else {
//		return errors.New("读取文件失败")
//	}
//}
//func myFn() {
//	defer func() {
//		err := recover()
//		if err != nil {
//			fmt.Println("给管理员发邮件")
//		}
//	}()
//	err := readFile("xx.go")
//	if err != nil {
//		panic(err)
//	}
//}
//func main() {
//	fmt.Println(fn1(10, 0))
//	myFn()
//}
