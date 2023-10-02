package main

//import "fmt"
//
//func printArray[T string | int](arr []T) {
//	for _, v := range arr {
//		fmt.Println(v)
//	}
//}
//
//func main() {
//	is := []int{1, 2}
//	strs := []string{"a", "b"}
//	printArray(is)
//	printArray(strs)
//
//	type Slice[T int | float32 | float64] []T
//
//	var a Slice[int] = []int{1, 2, 3}
//	fmt.Println(a)
//	var b Slice[float64] = []float64{1.00, 2.00, 4.00}
//	fmt.Println(b)
//	var c Slice[float32] = []float32{1.0, 2.0, 5.0}
//	fmt.Println(c)
//
//	type MyMap[KEY string | int, VALUE float32 | float64] map[KEY]VALUE
//
//	var m1 MyMap[string, float64] = map[string]float64{
//		"go":   9.9,
//		"java": 9.0,
//	}
//	fmt.Println(m1)
//}
