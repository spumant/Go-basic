package main

//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//	str1 := "this is str"
//	fmt.Println(str1)
//
//	str2 := `
//	this is str
//	this is str
//	`
//	fmt.Println(str2)
//
//	str3 := "aaa"
//	fmt.Println(len(str3))
//
//	str4 := "你好"
//	str5 := "golang"
//	fmt.Println(str4 + str5)
//
//	str6 := "123-456-789"
//	str7 := strings.Split(str6, "-")
//	fmt.Println(str7)
//
//	str8 := strings.Join(str7, "*")
//	fmt.Println(str8)
//
//	str9 := "this is str"
//	str10 := "this"
//	flag := strings.Contains(str9, str10)
//	fmt.Println(flag)
//
//	// 前缀后缀判断
//	flag1 := strings.HasPrefix(str9, str10) // 前缀
//	fmt.Println(flag1)
//
//	flag2 := strings.HasSuffix(str9, str10) // 后缀
//	fmt.Println(flag2)
//
//	// 查找子串，查找不到返回-1
//	str11 := "is"
//	num := strings.Index(str9, str11)
//	fmt.Println(num)
//
//	num2 := strings.LastIndex(str9, str11)
//	fmt.Println(num2)
//
//	s := "hello 张三"
//	for i := 0; i < len(s); i++ {
//		fmt.Printf("%v(%c)", s[i], s[i])
//	}
//	fmt.Println()
//	for _, r := range s {
//		fmt.Printf("%v(%c)", r, r)
//	}
//	fmt.Println()
//
//	changeStr()
//
//}
//func changeStr() {
//	s1 := "big"
//	// 强制类型转换
//	byteS1 := []byte(s1)
//	byteS1[0] = 'p'
//	fmt.Println(string(byteS1))
//
//	s2 := "白萝卜"
//	runeS2 := []rune(s2)
//	runeS2[0] = '红'
//	fmt.Println(string(runeS2))
//}
