package main

//import (
//	"fmt"
//	"sort"
//)
//
//type testSlice [][]int
//
//func (l testSlice) Len() int {
//	return len(l)
//}
//func (l testSlice) Swap(i, j int) {
//	l[i], l[j] = l[j], l[i]
//}
//
//func (l testSlice) Less(i, j int) bool {
//	return l[i][1] < l[j][1]
//}
//
//type testMap []map[string]float64
//
//func (m testMap) Len() int {
//	return len(testMap{})
//}
//
//func (m testMap) Swap(i, j int) {
//	m[i], m[j] = m[j], m[i]
//}
//
//func (m testMap) Less(i, j int) bool {
//	return m[i]["a"] < m[j]["a"]
//}
//
//type Person struct {
//	Name string
//	Age  int
//}
//
//type testStruct []Person
//
//func (u testStruct) Len() int {
//	return len(u)
//}
//
//func (u testStruct) Swap(i, j int) {
//	u[i], u[j] = u[j], u[i]
//}
//
//func (u testStruct) Less(i, j int) bool {
//	return u[i].Age < u[j].Age
//}
//
//func main() {
//	s := []int{2, 3, 1, 4}
//	sort.Ints(s)
//	fmt.Println(s)
//
//	ls := testSlice{
//		{1, 4},
//		{9, 3},
//		{7, 5},
//	}
//	sort.Sort(ls)
//	fmt.Println(ls)
//
//	lm := testMap{
//		{"a": 4, "b": 12},
//		{"a": 3, "b": 11},
//		{"a": 5, "b": 10},
//	}
//	sort.Sort(lm)
//	fmt.Println(lm)
//
//	lu := testStruct{
//		{Name: "n1", Age: 12},
//		{Name: "n2", Age: 11},
//		{Name: "n3", Age: 10},
//	}
//	sort.Sort(lu)
//	fmt.Println(lu)
//}
