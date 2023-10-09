package main

import (
	"Go-basic/12-Protobuf/proto/userService"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	u := &userService.Userinfo{
		Username: "张三",
		Age:      20,
		Hobby:    []string{"eat", "sleep"},
	}
	fmt.Println(u)

	fmt.Println(u.GetUsername())
	// 序列化Protobuf
	data, _ := proto.Marshal(u)
	fmt.Println(data)
	// 反序列化
	user := userService.Userinfo{}
	proto.Unmarshal(data, &user)
	fmt.Println(user)
	fmt.Println(user.GetHobby())
}
