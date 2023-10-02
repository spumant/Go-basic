package main

//import (
//	"context"
//	"fmt"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"log"
//)
//
//type Student struct {
//	Name string
//	Age  int
//}
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
//func insert(s Student) {
//	initDB()
//	collection := client.Database("go_db").Collection("student")
//	one, err := collection.InsertOne(context.TODO(), s)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(one.InsertedID)
//}
//
//func insertMany(students []interface{}) {
//	initDB()
//	collection := client.Database("go_db").Collection("student")
//	many, err := collection.InsertMany(context.TODO(), students)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(many.InsertedIDs)
//}
//
//func main() {
//	s := Student{
//		Name: "tom",
//		Age:  20,
//	}
//	insert(s)
//	s1 := Student{
//		Name: "kite",
//		Age:  18,
//	}
//	s2 := Student{
//		Name: "rose",
//		Age:  21,
//	}
//
//	stu := []interface{}{s1, s2}
//	insertMany(stu)
//}
