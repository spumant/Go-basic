package main

//import (
//	"errors"
//	"fmt"
//	"time"
//)
//
//func check(s string) (string, error) {
//	if s == "" {
//		err := errors.New("字符串不能为空")
//		return "", err
//	} else {
//		return s, nil
//	}
//}
//
//type MyError struct {
//	When time.Time
//	What string
//}
//
//func (e MyError) Error() string {
//	return fmt.Sprintf("%v : %v", e.When, e.What)
//}
//
//func oops() error {
//	return MyError{
//		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
//		"the file system has gone away",
//	}
//}
//
//func main() {
//	//s, err := check("")
//	//if err != nil {
//	//	fmt.Println(err)
//	//} else {
//	//	fmt.Println(s)
//	//}
//	err := oops()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//}
