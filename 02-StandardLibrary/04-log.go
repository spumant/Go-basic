package main

//import (
//	"fmt"
//	"log"
//	"os"
//)
//
//var Logger *log.Logger
//
//func test1() {
//	defer fmt.Println("panic 结束后再执行")
//	log.Println("my log")
//	log.Panic("my panic")
//	fmt.Println("end...")
//}
//
//func test2() {
//	defer fmt.Println("defrt...")
//	log.Println("my log")
//	log.Fatal("fatal...")
//	fmt.Println("end...")
//}
//
//func init() {
//	/*log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
//	log.SetPrefix("MyLog: ")
//	file, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
//	if err != nil {
//		log.Fatal("日志文件错误")
//	}
//	log.SetOutput(file)*/
//
//	file, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
//	if err != nil {
//		log.Fatal("日志文件错误")
//	}
//	Logger = log.New(file, "MyLog", log.Ldate|log.Ltime|log.Lshortfile)
//}
//func main() {
//	//i := log.Flags()
//	//fmt.Println(i)
//	log.Println("my log")
//}
