package main

//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"log"
//	"time"
//)
//
//var db *gorm.DB
//
//func init() {
//	dsn := "root:monarch99@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=True"
//	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	db = d
//}
//
//type User struct {
//	gorm.Model
//	Name     string
//	Age      int
//	Birthday time.Time
//}
//
//func createTable() {
//	db.AutoMigrate(&User{})
//}
//
//func insert1() {
//	user := User{
//		Name:     "tom",
//		Age:      20,
//		Birthday: time.Now(),
//	}
//	result := db.Create(&user)
//	fmt.Println(result.RowsAffected)
//	fmt.Println(user.ID)
//}
//
//func insert2() {
//	user := User{
//		Name: "mike",
//		Age:  18,
//	}
//	db.Select("Name", "Age", "CreatedAt").Create(&user)
//}
//
//func insert3() {
//	var users = []User{{Name: "x1", Birthday: time.Now()}, {Name: "x2", Birthday: time.Now()}, {Name: "x3", Birthday: time.Now()}}
//	db.Create(&users)
//}
//
//func main() {
//	//createTable()
//	//insert1()
//	//insert2()
//	insert3()
//}
