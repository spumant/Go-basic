package main

import (
	"fmt"
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

func testRaw1() {
	type Result struct {
		ID   int
		Name string
		Age  int
	}
	var result Result
	db.Raw("select id,name,age from users where name=?", "tom").Scan(&result)
	fmt.Println(result)

	var age int
	db.Raw("select sum(age) from users").Scan(&age)
	fmt.Println(age)
}

func testRaw2() {
	db.Exec("update users set age=? where name=?", 100, "tom")
}

func testRaw3() {
	var name string
	var age int
	row := db.Table("users").Where("name = ?", "tom").Select("name", "age").Row()
	row.Scan(&name, &age)
	fmt.Println(name, age)
}

func main() {

}
