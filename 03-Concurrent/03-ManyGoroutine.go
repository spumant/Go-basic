package main

//import (
//	"fmt"
//	"sync"
//)
//
//var wg sync.WaitGroup
//
//func count(num int) {
//	defer wg.Done()
//	for i := 0; i < 10; i++ {
//		fmt.Printf("协程（%v）打印的第%v条数据\n", num, i)
//	}
//}
//
//func main() {
//	for i := 1; i <= 5; i++ {
//		wg.Add(1)
//		go count(i)
//	}
//	wg.Wait()
//	fmt.Println("关闭主线程")
//}
