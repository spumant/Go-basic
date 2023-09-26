package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println(math.Pi)
	modf, frac := math.Modf(3.14)
	fmt.Println(modf, "--", frac)
	fmt.Println(rand.Int())
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println()
	}
}
