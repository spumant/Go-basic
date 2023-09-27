package main

//import (
//	"fmt"
//	"sync"
//)
//
//var wg sync.WaitGroup
//
//// 写数据
//func fn1(ch chan int) {
//	defer wg.Done()
//	for i := 0; i < 10; i++ {
//		ch <- i
//	}
//	close(ch)
//}
//
//// 读数据
//func fn2(ch chan int) {
//	defer wg.Done()
//	for v := range ch {
//		fmt.Println(v)
//	}
//}
//
//func main() {
//	var ch = make(chan int, 10)
//	wg.Add(1)
//	go fn1(ch)
//	wg.Add(1)
//	go fn2(ch)
//
//	wg.Wait()
//	fmt.Println("end...")
//}
