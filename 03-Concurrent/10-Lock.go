package main

//import (
//	"fmt"
//	"sync"
//)
//
//var wg sync.WaitGroup
//var mutex sync.RWMutex
//
//func Write() {
//	mutex.Lock()
//	fmt.Println("执行写操作")
//	mutex.Unlock()
//	wg.Done()
//}
//
//func Read() {
//	mutex.RLock()
//	fmt.Println("执行读操作")
//	mutex.RUnlock()
//	wg.Done()
//}
//
//func main() {
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go Read()
//	}
//
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go Write()
//	}
//	wg.Wait()
//}
