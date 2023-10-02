package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client

type Student struct {
	Name string
	Age  int
}

func initDB() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to MongoDB")
}

func delete() {
	ctx := context.TODO()
	defer client.Disconnect(ctx)
	c := client.Database("go_db").Collection("student")

	many, err := c.DeleteMany(ctx, bson.D{{"name", "big tom"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(many.DeletedCount)
}

func main() {
	initDB()
	delete()
}
