package main

import "fmt"

type int8AAA int8

type MyInt interface {
	int | ~int8 | int16 | int32 | int64
}

func GetMaxNum[T MyInt](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	var a int8AAA = 10
	var b int8AAA = 20
	c := GetMaxNum(a, b)
	fmt.Println(c)
}
