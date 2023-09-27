package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello world" + fmt.Sprintf("%d", i)
	}

	for {
		select {
		case v := <-intChan:
			fmt.Println(v)
		case v := <-stringChan:
			fmt.Println(v)

		}
	}
}
