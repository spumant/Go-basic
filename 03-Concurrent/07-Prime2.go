package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func putNum(intChan chan int) {
	for i := 2; i <= 120000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

func primeNum(intChan, primeChan chan int, exitChan chan bool) {
	for v := range intChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- v
		}
	}
	wg.Done()
	exitChan <- true
}

func printPrime(primeChan chan int) {
	for v := range primeChan {
		fmt.Println(v)
	}
	wg.Done()
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 16)

	wg.Add(1)
	go putNum(intChan)

	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}
	wg.Add(1)
	go printPrime(primeChan)

	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
}
