package main

//import (
//	"context"
//	"fmt"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"log"
//)
//
//var client *mongo.Client
//
//type Student struct {
//	Name string
//	Age  int
//}
//
//func initDB() {
//	// 设置客户端连接配置
//	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
//	// 连接到MongoDB
//	var err error
//	client, err = mongo.Connect(context.TODO(), clientOptions)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// 检查连接
//	err = client.Ping(context.TODO(), nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Connect to MongoDB")
//}
//
//func update() {
//	ctx := context.TODO()
//	defer client.Disconnect(ctx)
//	c := client.Database("go_db").Collection("student")
//
//	update := bson.D{{"$set", bson.D{{"name", "big tom"}, {"age", 22}}}}
//
//	many, err := c.UpdateMany(ctx, bson.D{{"name", "tom"}}, update)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(many.ModifiedCount)
//}
//
//func main() {
//	initDB()
//	update()
//}
