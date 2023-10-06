package main
//
//import (
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//	"log"
//	"time"
//)
//
//var db *gorm.DB
//
//func init() {
//	dsn := "root:monarch99@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=True"
//	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
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
//	Active   bool
//}
//
//func delete() {
//	var user User
//	db.First(&user)
//	db.Delete(&user)
//}
//
//func delete2() {
//	db.Delete(&User{}, 1)
//}
//
//func main() {
//
//}
