package main

//import (
//	"fmt"
//)
//
//func main() {
//	ch := make(chan int, 10)
//	ch <- 3
//	ch <- 12
//	ch <- 20
//	a := <-ch
//	fmt.Println(a)
//	<-ch
//	c := <-ch
//	fmt.Println(c)
//	// 值、容量、长度
//	fmt.Println(ch, cap(ch), len(ch))
//	for i := 1; i <= 10; i++ {
//		ch <- i
//	}
//	close(ch)  // 使用range遍历管道需要关闭，但通过for循环不需要
//	for val := range ch {
//		fmt.Println(val)
//	}
//
//}
