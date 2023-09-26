package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func test() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}
	marshal, _ := json.Marshal(p)
	fmt.Println(string(marshal))
}

func test2() {
	b := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com"}`)
	var m Person
	json.Unmarshal(b, &m)
	fmt.Println(m)
}

func test3() {
	b := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com","Parents":["big tom","kite"]}`)
	var p interface{}
	json.Unmarshal(b, &p)
	fmt.Println(p)
}

func test4() {
	open, _ := os.Open("a.json")
	defer open.Close()
	d := json.NewDecoder(open)
	var v map[string]interface{}
	d.Decode(&v)
	fmt.Println(v)

	for k, v := range v {
		fmt.Println(k, "--", v)
	}
}

func main() {
	test2()
}
