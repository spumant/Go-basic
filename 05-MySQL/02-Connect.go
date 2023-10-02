package main

//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//)
//
//var db *sql.DB
//
//func initDB() (err error) {
//	dsn := "root:monarch99@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=True"
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//	// 尝试与数据库建立连接
//	err = db.Ping()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func main() {
//	err := initDB()
//	if err != nil {
//		fmt.Println("连接失败", err)
//		return
//	} else {
//		fmt.Println("连接成功")
//	}
//}
