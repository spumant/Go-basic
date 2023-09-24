package main

import "fmt"

type Usber interface {
	Start()
	Stop()
}
type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "开始工作")
}
func (p Phone) Stop() {
	fmt.Println("phone 停止")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机 开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机 停止工作")
}

type Computer struct {
	Name string
}

func (c Computer) Work(usb Usber) {
	usb.Start()
	usb.Stop()
}
func main() {
	phone := Phone{Name: "小米手机"}
	var p Usber = phone
	p.Start()

	carema := Camera{}
	var c Usber = carema
	c.Start()

	computer := Computer{}
	computer.Work(phone)
	computer.Work(carema)

	// 定义一个空接口 x, x 变量可以接收任意的数据类型
	var x interface{}
	s := "你好 golang"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)

	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "张三"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

	var slice = []interface{}{"张三", 20, true, 32.2}
	fmt.Println(slice)

	// 类型断言
	var y interface{}
	y = "hello"
	v, ok := y.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
