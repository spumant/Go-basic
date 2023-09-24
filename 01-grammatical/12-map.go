package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//scoreMap := make(map[string]int, 8)
	//scoreMap["小明"] = 90
	//scoreMap["张三"] = 100
	//fmt.Println(scoreMap)
	//
	//userInfo := map[string]string{
	//	"username": "admin",
	//	"password": "123456",
	//}
	//fmt.Println(userInfo)
	//
	//// 判断某个键是否存在
	//v, ok := scoreMap["张三"]
	//if ok {
	//	fmt.Println(v)
	//} else {
	//	fmt.Println("查无此人")
	//}
	//
	//for k, v := range scoreMap {
	//	fmt.Println(k, "--", v)
	//}
	//
	//for k := range scoreMap {
	//	fmt.Println(k)
	//}
	//
	//delete(scoreMap, "小明")
	//for k, v := range scoreMap {
	//	fmt.Println(k, "--", v)
	//}

	// 按照指定顺序遍历map
	rand.Seed(time.Now().UnixMicro())
	scoreMap := make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		val := rand.Intn(100)
		scoreMap[key] = val
	}
	keys := make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	var mapSlice = make([]map[string]string, 3)
	for index, val := range mapSlice {
		fmt.Println(index, "--", val)
	}

	fmt.Println("----------------------------------")

	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "海淀区"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
