package main

//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//var wg sync.WaitGroup
//
//func test(n int) {
//	for num := (n-1)*30000 + 1; num < n*30000; num++ {
//		if num > 1 {
//			var flag = true
//			for i := 2; i < num; i++ {
//				if num%i == 0 {
//					flag = false
//					break
//				}
//			}
//			if flag {
//				//fmt.Println(num, "是素数")
//			}
//		}
//
//	}
//	wg.Done()
//}
//
//func main() {
//	start := time.Now().Unix()
//	for i := 1; i <= 4; i++ {
//		wg.Add(1)
//		go test(i)
//	}
//	wg.Wait()
//	end := time.Now().Unix()
//	fmt.Println(end - start)
//}
