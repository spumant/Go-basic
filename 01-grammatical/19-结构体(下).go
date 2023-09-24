package main

import "fmt"

type Person struct {
	name string
	age  int
}

// 值类型接受者
func (p Person) printInfo() {
	fmt.Printf("姓名:%v 年龄:%v", p.name, p.age)
}

// 指针类型接收者
func (p *Person) setInfo(name string, age int) {
	p.name = name
	p.age = age
}

type Address struct {
	Province string
	City     string
}

type User struct {
	Name    string
	Gender  string
	Address Address
}

type User2 struct {
	Name   string
	Gender string
	Address
}

type Animal struct {
	name string
}

func (a *Animal) run() {
	fmt.Printf("%s会运动", a.name)
}

type Dog struct {
	Age int
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s 会汪汪汪~\n", d.name)
}

func main() {
	p1 := Person{
		name: "小王子",
		age:  20,
	}
	p1.printInfo()
	p1.setInfo("张三", 18)
	p1.printInfo()

	user1 := User{
		Name:   "张三",
		Gender: "男",
		Address: Address{
			Province: "广东",
			City:     "深圳",
		},
	}
	fmt.Println(user1)

	var user2 User2
	user2.Name = "张三"
	user2.Gender = "男"
	user2.Address.Province = "广东"    //通过匿名结构体.字段名访问
	user2.City = "深圳"                //直接访问匿名结构体的字段名
	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:" 张 三 ", Gender:" 男 ", Address:main.Address{Province:"广东", City:"深圳"}}

	d1 := &Dog{
		Age: 4,
		Animal: &Animal{
			name: "阿奇",
		},
	}
	d1.wang()
	d1.run()
}
