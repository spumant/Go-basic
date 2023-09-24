package main

import "fmt"

type AInterface interface {
	GetInfo() string
}

type BInterface interface {
	SetInfo(string, int)
}
type People struct {
	Name string
	Age  int
}

func (p People) GetInfo() string {
	return fmt.Sprintf("姓名:%v 年龄:%d", p.Name, p.Age)
}
func (p *People) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}

type SayInterface interface {
	say()
}
type MoveInterface interface {
	move()
}

// 接口嵌套
type Animal interface {
	SayInterface
	MoveInterface
}
type Cat struct {
	name string
}

func (c Cat) say() {
	fmt.Println("喵喵喵")
}
func (c Cat) move() {
	fmt.Println("猫会动")
}
func main() {
	var people = &People{
		Name: "张三",
		Age:  20,
	}
	// people 实现了 AInterface 和 BInterface
	var p1 AInterface = people
	var p2 BInterface = people
	fmt.Println(p1.GetInfo())
	p2.SetInfo("李四", 30)
	fmt.Println(p1.GetInfo())

	var x Animal
	x = Cat{name: "花花"}
	x.move()
	x.say()
}
