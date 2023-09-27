package main

//import "fmt"
//
///*
//结构体首字母可以大写也可以小写，大写表示这个结构体是公有的，在其他的包里面
//可以使用。小写表示这个结构体是私有的，只有这个包里面才能使用。
//*/
//type person struct {
//	name string
//	city string
//	age  int
//}
//
//func main() {
//	var p1 person
//	p1.name = "zhangsan"
//	p1.city = "Beijing"
//	p1.age = 18
//	fmt.Println(p1)
//
//	/*
//		在 Golang 中支持对结构体指针直接使用.来访问结构体的成员。p2.name = "张三" 其
//		实在底层是(*p2).name = "张三"
//	*/
//	var p2 = new(person)
//	p2.name = "张三"
//	p2.age = 20
//	p2.city = "北京"
//	fmt.Printf("%T\n", p2)     //*main.person
//	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"张三", city:"北京", age:20}
//
//	p3 := &person{}
//	fmt.Printf("%T\n", p3)     //*main.person
//	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
//	p3.name = "zhangsan"
//	p3.age = 30
//	p3.city = "深圳"
//	(*p3).age = 40             //这样也是可以的
//	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"zhangsan", city:"深圳", age:30}
//
//	p4 := person{
//		name: "zhangsan",
//		city: "北京",
//		age:  20,
//	}
//	fmt.Println(p4)
//
//	p5 := &person{
//		name: "zhangsan",
//		city: "上海",
//		age:  28,
//	}
//	fmt.Printf("p5=%#v\n", p5) //p5=&main.person{name:"zhangsan", city:"上海", age:28}
//
//	p6 := &person{
//		city: "北京",
//	}
//	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"", city:"北京", age:0}
//
//	p7 := &person{
//		"zhangsan",
//		"北京",
//		28,
//	}
//	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"zhangsan", city:"北京", age:28}
//
//
//}
