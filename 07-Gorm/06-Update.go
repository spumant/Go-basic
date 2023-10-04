package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var db *gorm.DB

func init() {
	dsn := "root:monarch99@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=True"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
	Active   bool
}

func createTable() {
	db.AutoMigrate(&User{})
}

func insert() {
	user := User{
		Name:     "tom",
		Age:      20,
		Birthday: time.Now(),
		Active:   true,
	}
	db.Create(&user)
}

func update() {
	var user User
	db.First(&user)

	user.Name = "big tom"
	user.Age = 100
	db.Save(&user)
}

func update2() {
	db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
}

func update3() {
	var user User
	db.First(&user)
	db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
}

func main() {
	//createTable()
	//insert()
	update()
}
