package main

//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"log"
//)
//
//type Product struct {
//	gorm.Model
//	Code  string
//	Price uint
//}
//
//// 创建表
//func create(db *gorm.DB) {
//	db.AutoMigrate(&Product{})
//}
//
//// 插入数据
//func insert(p Product, db *gorm.DB) {
//	db.Create(&p)
//}
//
//// 查询
//func findById(db *gorm.DB) {
//	var p Product
//	db.First(&p, 1) // 主键为1
//	fmt.Println(p)
//}
//
//func findByCode(db *gorm.DB) {
//	var p Product
//	db.First(&p, "code = ?", "1001")
//	fmt.Println(p)
//}
//
//func update(db *gorm.DB) {
//	var p Product
//	db.First(&p, 1)
//	db.Model(&p).Update("price", 200)
//}
//
//func updateMany(db *gorm.DB) {
//	var p Product
//	db.First(&p, 1)
//	db.Model(&p).Updates(Product{Price: 1000, Code: "1002"})
//}
//
//func delete(db *gorm.DB) {
//	var p Product
//	db.Delete(&p, 1)
//}
//
//func main() {
//	dsn := "root:monarch99@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=True"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//create(db)
//	//
//	//p := Product{
//	//	Code:  "1001",
//	//	Price: 100,
//	//}
//	//
//	//insert(p, db)
//	//findById(db)
//	//findByCode(db)
//	//update(db)
//	//findById(db)
//	//updateMany(db)
//	//findById(db)
//	delete(db)
//}
