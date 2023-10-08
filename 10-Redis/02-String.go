package main

//import (
//	"context"
//	"fmt"
//	"github.com/redis/go-redis/v9"
//	"time"
//)
//
//var rdb *redis.Client
//
//func init() {
//	rdb = redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "",
//		DB:       0,
//	})
//}
//func main() {
//	ctx := context.Background()
//	err := rdb.Set(ctx, "gorediskey", "goreisvalue", 10*time.Second).Err()
//	if err != nil {
//		panic(err)
//	}
//	//// 将值改变并返回旧值
//	oldValue, err := rdb.GetSet(ctx, "gorediskey", "newvalue").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(oldValue)
//	//如果key不存在，则设置这个key的值
//	err = rdb.SetNX(ctx, "key1", "value", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//}
