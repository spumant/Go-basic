package main

//import (
//	"context"
//	"fmt"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"log"
//	"time"
//)
//
//var client *mongo.Client
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
//func find() {
//	defer client.Disconnect(context.TODO())
//	timeout, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancelFunc()
//	collection := client.Database("go_db").Collection("student")
//	cur, err := collection.Find(timeout, bson.D{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer cur.Close(timeout)
//	for cur.Next(timeout) {
//		var result bson.D
//		err := cur.Decode(&result)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(result)
//		fmt.Println(result.Map()["name"])
//	}
//	if err := cur.Err(); err != nil {
//		log.Fatal(err)
//	}
//}
//
//func main() {
//	initDB()
//	find()
//}
