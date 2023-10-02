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
//func insertData() {
//	sqlStr := "insert into user_tbl(username,password) values (?,?)"
//	ret, err := db.Exec(sqlStr, "张三", "zs123")
//	if err != nil {
//		fmt.Println("insert failed:", err)
//		return
//	}
//	id, err := ret.LastInsertId()
//	if err != nil {
//		fmt.Println("get lastinsert ID failed:", err)
//		return
//	}
//	fmt.Println("insert success,the id is ", id)
//}
//
//func insertData2(username, password string) {
//	sqlStr := "insert into user_tbl(username,password) values (?,?)"
//	ret, err := db.Exec(sqlStr, username, password)
//	if err != nil {
//		fmt.Println("insert failed:", err)
//		return
//	}
//	id, err := ret.LastInsertId()
//	if err != nil {
//		fmt.Println("get lastinsert ID failed:", err)
//		return
//	}
//	fmt.Println("insert success,the id is ", id)
//}
//
//func main() {
//	err := initDB()
//	if err != nil {
//		fmt.Println("初始化失败", err)
//	} else {
//		fmt.Println("初始化成功")
//	}
//	insertData()
//}
