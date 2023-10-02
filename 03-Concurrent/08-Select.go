package main

//import "fmt"
//
//func main() {
//	intChan := make(chan int, 10)
//	for i := 0; i < 10; i++ {
//		intChan <- i
//	}
//
//	stringChan := make(chan string, 5)
//	for i := 0; i < 5; i++ {
//		stringChan <- "hello world" + fmt.Sprintf("%d", i)
//	}
//	// 使用select来获取channel中的数据不需要关闭
//	for {
//		select {
//		case v := <-intChan:
//			fmt.Println(v)
//		case v := <-stringChan:
//			fmt.Println(v)
//		default:
//			fmt.Println("数据获取完毕")
//			return
//		}
//	}
//}
