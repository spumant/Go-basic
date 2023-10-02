package main

//import "fmt"
//
//type MySlice[T int | float64] []T
//
//func (s MySlice[T]) Sum() T {
//	var sum T
//	for _, v := range s {
//		sum += v
//	}
//	return sum
//}
//
//func Add[T int | float64 | string](a, b T) T {
//	return a + b
//}
//
//func main() {
//	var s MySlice[int] = []int{1, 2, 3, 4}
//	fmt.Println(s.Sum())
//
//	var f MySlice[float64] = []float64{1.2, 2.3, 3.4, 4.5}
//	fmt.Println(f.Sum())
//
//	fmt.Println(Add[int](1, 2))
//	fmt.Println(Add[string]("a", "hhh"))
//
//
//}
