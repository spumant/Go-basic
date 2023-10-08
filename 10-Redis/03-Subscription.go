package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
func main() {
	ctx := context.Background()
	sub := rdb.Subscribe(ctx, "channel1")
	for ch := range sub.Channel() {
		fmt.Println(ch.Channel) // 通道名
		fmt.Println(ch.Payload) // 值
	}
	// 或者
	for {
		message, err := sub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(message.Channel)
		fmt.Println(message.Payload)
	}
	// 取消订阅
	sub.Unsubscribe(ctx, "channel1")

	// 可以匹配ch_user_开头的任意channel
	sub2 := rdb.PSubscribe(ctx, "ch_user_*")
	for ch := range sub2.Channel() {
		fmt.Println(ch.Channel) // 通道名
		fmt.Println(ch.Payload) // 值
	}

	// 查询channel_1通道的订阅者数量
	chs, _ := rdb.PubSubNumSub(ctx, "channel_1").Result()

	for ch, count := range chs {
		fmt.Println(ch)    // channel名字
		fmt.Println(count) // channel的订阅者数量
	}

}
