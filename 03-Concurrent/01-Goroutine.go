package main

//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//var wg sync.WaitGroup
//
//func test1() {
//	for i := 0; i < 10; i++ {
//		fmt.Println("test1() hello", i)
//		time.Sleep(time.Millisecond * 50)
//	}
//	wg.Done()  // 协程计数器-1
//}
//
//func test2() {
//	for i := 0; i < 10; i++ {
//		fmt.Println("test2() hello", i)
//		time.Sleep(time.Millisecond * 50)
//	}
//	wg.Done()  // 协程计数器-1
//}
//
//func main() {
//	wg.Add(1) // 协程计数器+1
//	go test1()
//	wg.Add(1)
//	go test2()
//	for i := 0; i < 10; i++ {
//		fmt.Println("main() hello", i)
//		time.Sleep(time.Millisecond * 50)
//	}
//	wg.Wait()  // 等待携程执行完毕
//	fmt.Println("主线程退出")
//}
