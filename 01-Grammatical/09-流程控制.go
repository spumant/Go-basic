package main

//func main() {
//	// if 特殊写法
//	if score := 59; score >= 90 {
//		fmt.Println("A")
//	} else if score > 75 {
//		fmt.Println("B")
//	} else {
//		fmt.Println("C")
//	}
//
//	// 循环
//	for i := 0; i < 10; i++ {
//		fmt.Print(i)
//	}
//	fmt.Println()
//	i := 0
//	for i < 10 {
//		fmt.Print(i)
//		i++
//	}
//	// 无限循环
//	k := 1
//	for {
//		if k <= 10 {
//			fmt.Println("ok~~", k)
//		} else {
//			break
//		}
//		k++
//	}
//
//	//键值循环
//	str := "abc上海"
//	for index, val := range str {
//		fmt.Printf("%d-%c\n", index, val)
//	}
//	count := utf8.RuneCountInString(str)
//	fmt.Println(count)
//	switchDemo5()
//
//	// 多重循环，使用label标出想break的循环
//label:
//	for i := 0; i < 2; i++ {
//		for j := 0; j < 10; j++ {
//			if j == 2 {
//				break label
//			}
//			fmt.Println(i, "-", j)
//		}
//	}
//	// 标出想continue的循环
//here:
//	for i := 0; i < 2; i++ {
//		for j := 0; j < 4; j++ {
//			if j == 2 {
//				continue here
//			}
//			fmt.Println("i j 的值", i, "-", j)
//		}
//	}
//
//	for i := 0; i < 10; i++ {
//		for j := 0; j < 10; j++ {
//			if j == 2 {
//				// 设置退出标签
//				goto breakTag
//			}
//			fmt.Printf("%v-%v\n", i, j)
//		}
//	}
//	return
//	// 标签
//breakTag:
//	fmt.Println("结束 for 循环")
//}
//func switchDemo5() {
//	s := "a"
//	switch s {
//	case "a":
//		fmt.Println("a")
//		fallthrough
//	case "b":
//		fmt.Println("b")
//	case "c":
//		fmt.Println("c")
//	default:
//		fmt.Println("...")
//	}
//}
